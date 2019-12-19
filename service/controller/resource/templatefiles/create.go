package templatefiles

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/key"
)

const (
	filesdir = "/files"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	cc.Status.Files, err = key.Render(cc.Spec, key.FilePath)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
