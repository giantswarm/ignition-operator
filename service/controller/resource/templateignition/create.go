package templateignition

import (
	"context"

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

	master, err := key.Render(cc, key.MasterTemplatePath, false)
	if err != nil {
		return microerror.Mask(err)
	}

	worker, err := key.Render(cc, key.WorkerTemplatePath, false)
	if err != nil {
		return microerror.Mask(err)
	}

	s := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: key.StatusConfigMapName(cr.Spec.ClusterID),
			Labels: map[string]string{
				"cluster":       cr.Spec.ClusterID,
				label.ManagedBy: project.Name(),
			},
		},
		StringData: map[string]string{ // Data?
			"master": master["."],
			"worker": worker["."],
		},
	}

	actualSecret, err := r.k8sClient.K8sClient().CoreV1().Secrets(key.DefaultNamespace).Update(s)
	if apierrors.IsNotFound(err) {
		actualSecret, err = r.k8sClient.K8sClient().CoreV1().Secrets(key.DefaultNamespace).Create(s)
		if err != nil {
			return microerror.Mask(err)
		}
	} else if err != nil {
		return microerror.Mask(err)
	}

	cr.Status.Secret = v1alpha1.IgnitionStatusSecret{
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
