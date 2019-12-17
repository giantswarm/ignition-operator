package templateunits

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/shurcooL/httpfs/vfsutil"

	"github.com/giantswarm/ignition-operator/data"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	b, err := vfsutil.ReadFile(data.Assets, "worker_template.yaml")
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, string(b))
	return nil
}
