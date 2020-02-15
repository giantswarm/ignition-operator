package templateignition

import (
	"context"
	"crypto/sha512"
	"encoding/hex"

	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/ignition-operator/pkg/label"
	"github.com/giantswarm/ignition-operator/pkg/project"
	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/key"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cr, err := key.ToIgnition(obj)
	if err != nil {
		return microerror.Mask(err)
	}
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	var renderedIgnition string
	{
		templatePath := key.WorkerTemplatePath
		if cr.Spec.IsMaster {
			templatePath = key.MasterTemplatePath
		}
		result, err := key.Render(cc, templatePath, false)
		if err != nil {
			return microerror.Mask(err)
		}
		renderedIgnition = result["."]
	}

	s := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: key.StatusSecretName(cr.Spec.ClusterID, cr.GetName()),
			Labels: map[string]string{
				"cluster.x-k8s.io/cluster-name": cr.Spec.ClusterID,
				label.ManagedBy:                 project.Name(),
			},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(&cr, cr.GroupVersionKind()),
			},
		},
		StringData: map[string]string{
			"value": renderedIgnition,
		},
	}

	actualSecret, err := r.k8sClient.K8sClient().CoreV1().Secrets(cr.GetNamespace()).Update(s)
	if apierrors.IsNotFound(err) {
		actualSecret, err = r.k8sClient.K8sClient().CoreV1().Secrets(cr.GetNamespace()).Create(s)
		if err != nil {
			return microerror.Mask(err)
		}
	} else if err != nil {
		return microerror.Mask(err)
	}

	cr.Status.Verification.Hash = calculateHash([]byte(renderedIgnition))
	cr.Status.Verification.Algorithm = "sha512"
	cr.Status.Ready = true
	cr.Status.DataSecret = v1alpha1.IgnitionStatusSecret{
		Name:            actualSecret.Name,
		Namespace:       actualSecret.Namespace,
		ResourceVersion: actualSecret.ResourceVersion,
	}

	_, err = r.k8sClient.G8sClient().CoreV1alpha1().Ignitions(cr.Namespace).UpdateStatus(&cr)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func calculateHash(data []byte) string {
	rawSum := sha512.Sum512(data)
	sum := rawSum[:]

	encodedSum := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(encodedSum, sum)

	return string(encodedSum)
}
