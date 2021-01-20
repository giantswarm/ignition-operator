package templateignition

import (
	"context"

	"github.com/giantswarm/apiextensions/v3/pkg/apis/core/v1alpha1"
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

	var template map[string]string
	{
		if cr.Spec.IsMaster {
			template, err = key.Render(cc, key.MasterTemplatePath, false)
		} else {
			template, err = key.Render(cc, key.WorkerTemplatePath, false)
		}
		if err != nil {
			return microerror.Mask(err)
		}
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
			"value": template["."],
		},
	}

	actualSecret, err := r.k8sClient.K8sClient().CoreV1().Secrets(cr.GetNamespace()).Update(ctx, s, metav1.UpdateOptions{})
	if apierrors.IsNotFound(err) {
		actualSecret, err = r.k8sClient.K8sClient().CoreV1().Secrets(cr.GetNamespace()).Create(ctx, s, metav1.CreateOptions{})
		if err != nil {
			return microerror.Mask(err)
		}
	} else if err != nil {
		return microerror.Mask(err)
	}

	cr.Status.DataSecret = v1alpha1.IgnitionStatusSecret{
		Name:            actualSecret.Name,
		Namespace:       actualSecret.Namespace,
		ResourceVersion: actualSecret.ResourceVersion,
	}

	_, err = r.k8sClient.G8sClient().CoreV1alpha1().Ignitions(cr.Namespace).UpdateStatus(ctx, &cr, metav1.UpdateOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
