package encryptionkey

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/key"
	"github.com/giantswarm/microerror"
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

	cc.Spec = crToCC(cr)

	return nil
}

func crToCC(cr v1alpha1.Ignition) controllercontext.ContextSpec {
	var ccSpec controllercontext.ContextSpec
	{
		ccSpec = controllercontext.ContextSpec{
			BaseDomain: cr.Spec.BaseDomain,
			Calico: controllercontext.ContextSpecCalico{
				CIDR:    cr.Spec.Calico.CIDR,
				Disable: cr.Spec.Calico.Disable,
				MTU:     cr.Spec.Calico.MTU,
				Subnet:  cr.Spec.Calico.Subnet,
			},
			Etcd: controllercontext.ContextSpecEtcd{
				Domain: cr.Spec.Etcd.Domain,
				Port:   cr.Spec.Etcd.Port,
				Prefix: cr.Spec.Etcd.Prefix,
			},
			Ingress: controllercontext.ContextSpecIngress{
				Disable: cr.Spec.Ingress.Disable,
			},
			Kubernetes: controllercontext.ContextSpecKubernetes{
				API: controllercontext.ContextSpecKubernetesAPI{
					Domain:     cr.Spec.Kubernetes.API.Domain,
					SecurePort: cr.Spec.Kubernetes.API.SecurePort,
				},
				DNS: controllercontext.ContextSpecKubernetesDNS{
					IP: cr.Spec.Kubernetes.DNS.IP,
				},
				Domain: cr.Spec.Kubernetes.Domain,
				Kubelet: controllercontext.ContextSpecKubernetesKubelet{
					Domain: cr.Spec.Kubernetes.Kubelet.Domain,
				},
				Image:   cr.Spec.Kubernetes.Image,
				IPRange: cr.Spec.Kubernetes.IPRange,
			},
			Provider: cr.Spec.Provider,
			Registry: controllercontext.ContextSpecRegistry{
				Domain: cr.Spec.Registry.Domain,
			},
			SSO: controllercontext.ContextSpecSSO{
				PublicKey: cr.Spec.SSO.PublicKey,
			},
		}
	}

	return ccSpec
}
