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
			DisableEncryptionAtRest: false,
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
			Extension: controllercontext.ContextSpecExtension{
				Files: []controllercontext.ContextSpecExtensionFile{
					{
						Content: "someFileContent",
						Metadata: controllercontext.ContextSpecExtensionFileMetadata{
							Compression: true,
							Owner: controllercontext.ContextSpecExtensionFileMetadataOwner{
								Group: controllercontext.ContextSpecExtensionFileMetadataOwnerGroup{
									ID:   "someGroupID",
									Name: "someGroupName",
								},
								User: controllercontext.ContextSpecExtensionFileMetadataOwnerUser{
									ID:   "someUserID",
									Name: "someUSerName",
								},
							},
							Path:        "some/File/Path",
							Permissions: 700,
						},
					},
					{
						Content: "someOtherFileContent",
						Metadata: controllercontext.ContextSpecExtensionFileMetadata{
							Compression: false,
							Owner: controllercontext.ContextSpecExtensionFileMetadataOwner{
								Group: controllercontext.ContextSpecExtensionFileMetadataOwnerGroup{
									ID:   "someOtherGroupID",
									Name: "someOtherGroupName",
								},
								User: controllercontext.ContextSpecExtensionFileMetadataOwnerUser{
									ID:   "someOtherUserID",
									Name: "someOtherUserName",
								},
							},
							Path:        "some/Other/File/Path",
							Permissions: 700,
						},
					},
				},
				Units: []controllercontext.ContextSpecExtensionUnit{
					{
						Content: `[Unit]
Description=Some sample Unit
After=network.target
[Service]
Type=oneshot
ExecStart=/opt/some-debug-unit
[Install]
WantedBy=multi-user.target`,
						Metadata: controllercontext.ContextSpecExtensionUnitMetadata{
							Enabled: true,
							Name:    "SomeUnit",
						},
					},
					{
						Content: `[Unit]
Description=Some other sample Unit
After=network.target
[Service]
Type=oneshot
ExecStart=/opt/some-debug-unit
[Install]
WantedBy=multi-user.target`,
						Metadata: controllercontext.ContextSpecExtensionUnitMetadata{
							Enabled: false,
							Name:    "SomeOtherUnit",
						},
					},
				},
				Users: []controllercontext.ContextSpecExtensionUser{
					{
						Name:      "SomeUser",
						PublicKey: "SomePubKey",
					},
					{
						Name: "UserWithoutPubKey",
					},
				},
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
