[![CircleCI](https://circleci.com/gh/giantswarm/ignition-operator.svg?&style=shield)](https://circleci.com/gh/giantswarm/ignition-operator)
[![Docker Repository on Quay](https://quay.io/repository/giantswarm/ignition-operator/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/ignition-operator)

# ignition-operator

A [Kubernetes operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) based on
[Giant Swarm's operatorkit](https://github.com/giantswarm/operatorkit) which watches for
[Ignition resources](https://github.com/giantswarm/apiextensions/blob/cbcf4d4b80bf897536a4070b0689493d2780e143/pkg/apis/core/v1alpha1/ignition_types.go#L32),
applies node-specific data to [templates](https://github.com/giantswarm/ignition-operator/tree/master/template) using
the [go template library](https://golang.org/pkg/text/template/), and ultimately renders the full ignition into a secret
to be passed into nodes using [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on first boot.

## Code generation

After modifying template files in `template/`, you must regenerate the `vfsstatic.go` file using `go generate pkg/asset/generate.go`.
