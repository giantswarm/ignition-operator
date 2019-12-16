package key

import (
	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/ignition-operator/pkg/label"
	"github.com/giantswarm/microerror"
)

func OperatorVersion(getter LabelsGetter) string {
	return getter.GetLabels()[label.OperatorVersion]
}

func ToIgnition(v interface{}) (v1alpha1.Ignition, error) {
	if v == nil {
		return v1alpha1.Ignition{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.Ignition{}, v)
	}

	p, ok := v.(*v1alpha1.Ignition)
	if !ok {
		return v1alpha1.Ignition{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.Ignition{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}
