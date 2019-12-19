package controllercontext

type ContextSpec struct {
	BaseDomain              string
	Calico                  ContextSpecCalico
	DisableEncryptionAtREST bool
	Docker                  ContextSpecDocker
	Etcd                    ContextSpecEtcd
	Ingress                 ContextSpecIngress
	Kubernetes              ContextSpecKubernetes
	// Defines the provider which should be rendered.
	Provider string
	Registry ContextSpecRegistry
	SSO      ContextSpecSSO
}

type ContextSpecCalico struct {
	CIDR    string
	Disable bool
	MTU     string
	Subnet  string
}

type ContextSpecDocker struct {
	Daemon       ContextSpecDockerDaemon
	NetworkSetup ContextSpecDockerNetworkSetup
}

type ContextSpecDockerDaemon struct {
	CIDR string
}

type ContextSpecDockerNetworkSetup struct {
	Image string
}

type ContextSpecEtcd struct {
	Domain string
	Image  string
	Port   int
	Prefix string
}

type ContextSpecIngress struct {
	Disable bool
}

type ContextSpecKubernetes struct {
	API     ContextSpecKubernetesAPI
	DNS     ContextSpecKubernetesDNS
	Domain  string
	Kubelet ContextSpecKubernetesKubelet
	Image   string
	IPRange string
}

type ContextSpecKubernetesAPI struct {
	Domain     string
	SecurePort int
}

type ContextSpecKubernetesDNS struct {
	IP string
}

type ContextSpecKubernetesKubelet struct {
	CommandArgs []string
	Domain      string
	Labels      string
}

type ContextSpecRegistry struct {
	Domain               string
	PullProgressDeadline string
}
type ContextSpecSSO struct {
	PublicKey string
}
