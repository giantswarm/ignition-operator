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

	cc.Spec = crToCCSpec(cr)

	return nil
}

func crToCCSpec(cr v1alpha1.Ignition) controllercontext.ContextSpec {
	var ccExtensionFiles []controllercontext.ContextSpecExtensionFile
	{
		for _, crExtensionFile := range cr.Spec.Extension.Files {
			ccExtensionFile := controllercontext.ContextSpecExtensionFile{
				Content: crExtensionFile.Content,
				Metadata: controllercontext.ContextSpecExtensionFileMetadata{
					Compression: crExtensionFile.Metadata.Compression,
					Owner: controllercontext.ContextSpecExtensionFileMetadataOwner{
						Group: controllercontext.ContextSpecExtensionFileMetadataOwnerGroup{
							ID:   crExtensionFile.Metadata.Owner.Group.ID,
							Name: crExtensionFile.Metadata.Owner.Group.Name,
						},
						User: controllercontext.ContextSpecExtensionFileMetadataOwnerUser{
							ID:   crExtensionFile.Metadata.Owner.User.ID,
							Name: crExtensionFile.Metadata.Owner.User.Name,
						},
					},
					Path:        crExtensionFile.Metadata.Path,
					Permissions: crExtensionFile.Metadata.Permissions,
				},
			}
			ccExtensionFiles = append(ccExtensionFiles, ccExtensionFile)
		}
	}

	var ccExtensionUnits []controllercontext.ContextSpecExtensionUnit
	{
		for _, crExtensionUnit := range cr.Spec.Extension.Units {
			ccExtensionUnit := controllercontext.ContextSpecExtensionUnit{
				Content: crExtensionUnit.Content,
				Metadata: controllercontext.ContextSpecExtensionUnitMetadata{
					Enabled: crExtensionUnit.Metadata.Enabled,
					Name:    crExtensionUnit.Metadata.Name,
				},
			}
			ccExtensionUnits = append(ccExtensionUnits, ccExtensionUnit)
		}
	}

	var ccExtensionUsers []controllercontext.ContextSpecExtensionUser
	{
		for _, crExtensionUser := range cr.Spec.Extension.Users {
			ccExtensionUser := controllercontext.ContextSpecExtensionUser{
				Name:      crExtensionUser.Name,
				PublicKey: crExtensionUser.PublicKey,
			}
			ccExtensionUsers = append(ccExtensionUsers, ccExtensionUser)
		}
	}

	var ccSpec controllercontext.ContextSpec
	{
		ccSpec = controllercontext.ContextSpec{
			APIServerEncryptionKey: cr.Spec.APIServerEncryptionKey,
			BaseDomain:             cr.Spec.BaseDomain,
			Calico: controllercontext.ContextSpecCalico{
				CIDR:    cr.Spec.Calico.CIDR,
				Disable: cr.Spec.Calico.Disable,
				MTU:     cr.Spec.Calico.MTU,
				Subnet:  cr.Spec.Calico.Subnet,
			},
			DisableEncryptionAtREST: cr.Spec.DisableEncryptionAtREST,
			Docker: controllercontext.ContextSpecDocker{
				Daemon: controllercontext.ContextSpecDockerDaemon{
					CIDR: cr.Spec.Docker.Daemon.CIDR,
				},
				NetworkSetup: controllercontext.ContextSpecDockerNetworkSetup{
					Image: cr.Spec.Docker.NetworkSetup.Image,
				},
			},
			Etcd: controllercontext.ContextSpecEtcd{
				Domain: cr.Spec.Etcd.Domain,
				Image:  cr.Spec.Etcd.Image,
				Port:   cr.Spec.Etcd.Port,
				Prefix: cr.Spec.Etcd.Prefix,
			},
			Extension: controllercontext.ContextSpecExtension{
				Files: ccExtensionFiles,
				Units: ccExtensionUnits,
				Users: ccExtensionUsers,
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
				Domain:               cr.Spec.Registry.Domain,
				PullProgressDeadline: cr.Spec.Registry.PullProgressDeadline,
			},
			SSO: controllercontext.ContextSpecSSO{
				PublicKey: cr.Spec.SSO.PublicKey,
			},
		}
	}

	return ccSpec
}
