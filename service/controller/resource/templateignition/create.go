package templateignition

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/key"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	cc.Status.Units, err = key.Render(cc.Status, key.UnitPath, false)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
