package unittest

import "github.com/giantswarm/ignition-operator/service/controller/controllercontext"

func DefaultCC() controllercontext.Context {
	var ccSpec controllercontext.ContextSpec
	{
		ccSpec = controllercontext.ContextSpec{
			APIServerEncryptionKey: "some secret",
			BaseDomain:             "ClusterBaseDomain",
			Calico: controllercontext.ContextSpecCalico{
				CIDR:    "CalicoCIDR",
				Disable: false,
				MTU:     "CalicoMTU",
				Subnet:  "CalicoSubnet",
			},
			DisableEncryptionAtREST: false,
			Docker: controllercontext.ContextSpecDocker{
				Daemon: controllercontext.ContextSpecDockerDaemon{
					CIDR: "DockerDaemonCIDR",
				},
				NetworkSetup: controllercontext.ContextSpecDockerNetworkSetup{
					Image: "DockerNetworkSetupImage",
				},
			},
			Etcd: controllercontext.ContextSpecEtcd{
				Domain: "EtcdDomain",
				Image:  "EtcdImage",
				Port:   1234,
				Prefix: "EtcdPrefix",
			},
			Ingress: controllercontext.ContextSpecIngress{
				Disable: false,
			},
			Kubernetes: controllercontext.ContextSpecKubernetes{
				API: controllercontext.ContextSpecKubernetesAPI{
					Domain:     "APIDomain",
					SecurePort: 9001,
				},
				DNS: controllercontext.ContextSpecKubernetesDNS{
					IP: "K8sDNSIP",
				},
				Domain: "K8sDomain",
				Kubelet: controllercontext.ContextSpecKubernetesKubelet{
					CommandArgs: []string{
						"kubeletArg1",
						"kubeletArg2",
					},
					Domain: "K8sKubeletDomain",
					Labels: "some=label",
				},
				Image:   "K8sImage",
				IPRange: "K8sIPRange",
			},
			Provider: "aws",
			Registry: controllercontext.ContextSpecRegistry{
				Domain:               "RegistryDomain",
				PullProgressDeadline: "SomeProgressDeadline",
			},
			SSO: controllercontext.ContextSpecSSO{
				PublicKey: "SSOPublicKey",
			},
		}
	}

	cc := controllercontext.Context{
		Spec: ccSpec,
	}

	return cc
}
