package controller

import (
	"github.com/giantswarm/apiextensions/v3/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/v4/pkg/controller"
	"github.com/giantswarm/operatorkit/v4/pkg/resource"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/ignition-operator/pkg/project"
)

type IgnitionConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type Ignition struct {
	*controller.Controller
}

func NewIgnition(config IgnitionConfig) (*Ignition, error) {
	var err error

	resources, err := newIgnitionResources(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
			Resources: resources,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(v1alpha1.Ignition)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/ignition-operator-ignition-controller.
			Name: project.Name() + "-ignition-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &Ignition{
		Controller: operatorkitController,
	}

	return c, nil
}

func newIgnitionResources(config IgnitionConfig) ([]resource.Interface, error) {
	var err error

	resources := []resource.Interface{}

	var ignitionResources []resource.Interface
	{
		c := ignitionResourceSetConfig(config)

		ignitionResources, err = newIgnitionResource(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources = append(resources, ignitionResources...)

	return resources, nil
}
