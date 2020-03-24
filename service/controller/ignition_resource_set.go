package controller

import (
	"context"

	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"github.com/giantswarm/operatorkit/resource"
	"github.com/giantswarm/operatorkit/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/resource/wrapper/retryresource"

	"github.com/giantswarm/ignition-operator/pkg/project"
	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/key"
	"github.com/giantswarm/ignition-operator/service/controller/resource/contextspec"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templatefiles"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templateignition"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templateunits"
)

type ignitionResourceSetConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

func newIgnitionResourceSet(config ignitionResourceSetConfig) (*controller.ResourceSet, error) {
	var err error

	var contextspecResource resource.Interface
	{
		c := contextspec.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		contextspecResource, err = contextspec.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var templatefilesResource resource.Interface
	{
		c := templatefiles.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		templatefilesResource, err = templatefiles.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var templateunitsResource resource.Interface
	{
		c := templateunits.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		templateunitsResource, err = templateunits.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}
	var templateignitionResource resource.Interface
	{
		c := templateignition.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		templateignitionResource, err = templateignition.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources := []resource.Interface{
		contextspecResource,
		templatefilesResource,
		templateunitsResource,
		templateignitionResource,
	}

	{
		c := retryresource.WrapConfig{
			Logger: config.Logger,
		}

		resources, err = retryresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		c := metricsresource.WrapConfig{}

		resources, err = metricsresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	handlesFunc := func(obj interface{}) bool {
		cr, err := key.ToIgnition(obj)
		if err != nil {
			return false
		}

		if key.OperatorVersion(&cr) == project.Version() {
			return true
		}

		return false
	}

	initCtxFunc := func(ctx context.Context, obj interface{}) (context.Context, error) {
		return controllercontext.NewContext(ctx, controllercontext.Context{}), nil
	}

	var resourceSet *controller.ResourceSet
	{
		c := controller.ResourceSetConfig{
			Handles:   handlesFunc,
			InitCtx:   initCtxFunc,
			Logger:    config.Logger,
			Resources: resources,
		}

		resourceSet, err = controller.NewResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return resourceSet, nil
}
