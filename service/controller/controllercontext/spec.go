package controllercontext

type ContextSpec struct {
	APIServerEncryptionKey  string
	BaseDomain              string
	Calico                  ContextSpecCalico
	DisableEncryptionAtRest bool
	Docker                  ContextSpecDocker
	Etcd                    ContextSpecEtcd
	Extension               ContextSpecExtension
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

type ContextSpecExtension struct {
	Files []ContextSpecExtensionFile
	Units []ContextSpecExtensionUnit
	Users []ContextSpecExtensionUser
}

type ContextSpecExtensionFile struct {
	Content  string
	Metadata ContextSpecExtensionFileMetadata
}

type ContextSpecExtensionFileMetadata struct {
	Compression bool
	Owner       ContextSpecExtensionFileMetadataOwner
	Path        string
	Permissions int
}

type ContextSpecExtensionFileMetadataOwner struct {
	Group ContextSpecExtensionFileMetadataOwnerGroup
	User  ContextSpecExtensionFileMetadataOwnerUser
}

type ContextSpecExtensionFileMetadataOwnerUser struct {
	ID   string
	Name string
}

type ContextSpecExtensionFileMetadataOwnerGroup struct {
	ID   string
	Name string
}

type ContextSpecExtensionUnit struct {
	Content  string
	Metadata ContextSpecExtensionUnitMetadata
}

type ContextSpecExtensionUnitMetadata struct {
	Enabled bool
	Name    string
}

type ContextSpecExtensionUser struct {
	Name      string
	PublicKey string
}

type ContextSpecIngress struct {
	Disable bool
}

type ContextSpecKubernetes struct {
	API           ContextSpecKubernetesAPI
	CloudProvider string
	DNS           ContextSpecKubernetesDNS
	Domain        string
	Kubelet       ContextSpecKubernetesKubelet
	Image         string
	IPRange       string
	OIDC          ContextSpecOIDC
}

type ContextSpecOIDC struct {
	Enabled       bool
	ClientID      string
	IssuerURL     string
	UsernameClaim string
	GroupsClaim   string
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
