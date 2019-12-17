package test

import (
	"context"
	"io/ioutil"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/ignition-operator/data"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	rc, err := data.Assets.Open("master_template.yaml")
	if err != nil {
		return microerror.Mask(err)
	}
	defer rc.Close()

	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, string(b))
	return nil
}
