package controller

import (
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/v4/pkg/resource"
	"github.com/giantswarm/operatorkit/v4/pkg/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/v4/pkg/resource/wrapper/retryresource"

	"github.com/giantswarm/ignition-operator/service/controller/resource/contextspec"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templatefiles"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templateignition"
	"github.com/giantswarm/ignition-operator/service/controller/resource/templateunits"
)

type ignitionResourceSetConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

func newIgnitionResource(config ignitionResourceSetConfig) ([]resource.Interface, error) {
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

	return resources, nil
}
