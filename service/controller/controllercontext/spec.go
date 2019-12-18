package controllercontext

type ContextSpec struct {
	// APIServerEncryptionKey is AES-CBC with PKCS#7 padding key to encrypt API
	// etcd data.
	APIServerEncryptionKey string
	BaseDomain             string
	// DisableCalico flag. When set removes all calico related Kubernetes
	// manifests from the cloud config together with their initialization.
	DisableCalico bool
	// DisableEncryptionAtREST flag. When set removes all manifests from the cloud
	// config related to Kubernetes encryption at REST.
	DisableEncryptionAtREST bool
	// DisableIngressControllerService flag. When set removes the manifest for
	// the Ingress Controller service. This allows us to migrate providers to
	// chart-operator independently.
	DisableIngressControllerService bool
	// Hyperkube allows to pass extra `docker run` and `command` arguments
	// to hyperkube image commands. This allows to e.g. add cloud provider
	// extensions.
	Hyperkube Hyperkube
	// EtcdPort allows the Etcd port to be specified.
	// aws-operator sets this to the Etcd listening port so Calico on the
	// worker nodes can access via a CNAME record to the master.
	EtcdPort int
	// ImagePullProgressDeadline is the duration after which image pulling is
	// cancelled if no progress has been made.
	ImagePullProgressDeadline string
	// RegistryDomain is the host of the docker image registry to use.
	RegistryDomain string
	SSOPublicKey   string
	// Container images used in the cloud-config templates
	Images Images
}

type Images struct {
	Kubernetes string
	Etcd       string
}

type Hyperkube struct {
	Apiserver         HyperkubeApiserver
	ControllerManager HyperkubeControllerManager
	Kubelet           HyperkubeKubelet
}

type HyperkubeApiserver struct {
	Pod HyperkubePod
}

type HyperkubeControllerManager struct {
	Pod HyperkubePod
}

type HyperkubeKubelet struct {
	Docker HyperkubeDocker
}

type HyperkubeDocker struct {
	RunExtraArgs     []string
	CommandExtraArgs []string
}

type HyperkubePod struct {
	HyperkubePodHostExtraMounts []HyperkubePodHostMount
	CommandExtraArgs            []string
}

type HyperkubePodHostMount struct {
	Name     string
	Path     string
	ReadOnly bool
}
