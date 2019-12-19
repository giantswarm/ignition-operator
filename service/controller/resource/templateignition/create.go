package templateignition

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      key.StatusConfigMapName(cr.Spec.ClusterID),
			Namespace: key.DefaultNamespace,
		},
		Data: map[string]string{
			"master": master["."],
			"worker": worker["."],
		},
	}

	cm, err = r.k8sClient.K8sClient().CoreV1().ConfigMaps(key.DefaultNamespace).Update(cm)
	if apierrors.IsNotFound(err) {
		cm, err = r.k8sClient.K8sClient().CoreV1().ConfigMaps(key.DefaultNamespace).Create(cm)
		if err != nil {
			return microerror.Mask(err)
		}

	} else if err != nil {
		return microerror.Mask(err)
	}

	cr.Status.ConfigMap = v1alpha1.IgnitionStatusConfigMap{
		Name:            cm.Name,
		Namespace:       cm.Namespace,
		ResourceVersion: cm.ResourceVersion,
	}

	_, err = r.k8sClient.G8sClient().CoreV1alpha1().Ignitions(key.DefaultNamespace).UpdateStatus(&cr)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
