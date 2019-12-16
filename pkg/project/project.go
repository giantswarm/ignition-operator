package project

var (
	bundleVersion = "0.0.1"
	description   = "The ignition-operator does something."
	gitSHA        = "n/a"
	name          = "ignition-operator"
	source        = "https://github.com/giantswarm/ignition-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
