package project

var (
	description = "The ignition-operator does something."
	gitSHA      = "n/a"
	name        = "ignition-operator"
	source      = "https://github.com/giantswarm/ignition-operator"
	version     = "0.0.2-dev"
)

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
