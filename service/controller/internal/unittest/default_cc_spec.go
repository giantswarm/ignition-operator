package unittest

import "github.com/giantswarm/ignition-operator/service/controller/controllercontext"

func DefaultCCSpec() controllercontext.ContextSpec {
	var ccSpec controllercontext.ContextSpec
	{
		ccSpec = controllercontext.ContextSpec{
			BaseDomain: "ClusterBaseDomain",
			Calico: controllercontext.ContextSpecCalico{
				CIDR:    "CalicoCIDR",
				Disable: false,
				MTU:     "CalicoMTU",
				Subnet:  "CalicoSubnet",
			},
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

	return ccSpec
}
