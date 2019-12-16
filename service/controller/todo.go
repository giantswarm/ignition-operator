package controller

import (
	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/ignition-operator/pkg/project"
)

type TODOConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type TODO struct {
	*controller.Controller
}

func NewTODO(config TODOConfig) (*TODO, error) {
	var err error

	resourceSets, err := newTODOResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(v1alpha1.Ignition)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/ignition-operator-todo-controller.
			Name: project.Name() + "-todo-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &TODO{
		Controller: operatorkitController,
	}

	return c, nil
}

func newTODOResourceSets(config TODOConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := todoResourceSetConfig{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		resourceSet, err = newTODOResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
