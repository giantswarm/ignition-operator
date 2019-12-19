package controllercontext

type ContextSpec struct {
	BaseDomain string                `json:"basedomain" yaml:"basedomain"`
	Calico     ContextSpecCalico     `json:"calico" yaml:"calico"`
	Etcd       ContextSpecEtcd       `json:"etcd" yaml:"etcd"`
	Ingress    ContextSpecIngress    `json:"ingress" yaml:"ingress"`
	Kubernetes ContextSpecKubernetes `json:"kubernetes" yaml:"kubernetes"`
	// Defines the provider which should be rendered.
	Provider string              `json:"provider" yaml:"provider"`
	Registry ContextSpecRegistry `json:"registry" yaml:"registry"`
	SSO      ContextSpecSSO      `json:"sso" yaml:"sso"`
}

type ContextSpecCalico struct {
	CIDR    string `json:"cidr" yaml:"cidr"`
	Disable bool   `json:"disable" yaml:"disable"`
	MTU     string `json:"mtu" yaml:"mtu"`
	Subnet  string `json:"subnet" yaml:"subnet"`
}

type ContextSpecEtcd struct {
	Domain string `json:"domain" yaml:"domain"`
	Port   int    `json:"port" yaml:"port"`
	Prefix string `json:"prefix" yaml:"prefix"`
}

type ContextSpecIngress struct {
	Disable bool `json:"disable" yaml:"disable"`
}

type ContextSpecKubernetes struct {
	API     ContextSpecKubernetesAPI     `json:"api" yaml:"api"`
	DNS     ContextSpecKubernetesDNS     `json:"dns" yaml:"dns"`
	Domain  string                       `json:"domain" yaml:"domain"`
	Kubelet ContextSpecKubernetesKubelet `json:"kubelet" yaml:"kubelet"`
	Image   string                       `json:"image" yaml:"image"`
	IPRange string                       `json:"iprange" yaml:"iprange"`
}

type ContextSpecKubernetesAPI struct {
	Domain     string `json:"domain" yaml:"domain"`
	SecurePort int    `json:"secureport" yaml:"secureport"`
}

type ContextSpecKubernetesDNS struct {
	IP string `json:"ip" yaml:"ip"`
}

type ContextSpecKubernetesKubelet struct {
	Domain string `json:"domain" yaml:"domain"`
	Labels string `json:"labels" yaml:"labels"`
}

type ContextSpecRegistry struct {
	Domain string `json:"domain" yaml:"domain"`
}
type ContextSpecSSO struct {
	PublicKey string `json:"publicKey" yaml:"publicKey"`
}
