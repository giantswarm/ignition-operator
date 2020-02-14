[![CircleCI](https://circleci.com/gh/giantswarm/ignition-operator.svg?&style=shield)](https://circleci.com/gh/giantswarm/ignition-operator) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/ignition-operator/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/ignition-operator)

# ignition-operator

It templates ignition.

## Code generation

After modifying files in `data/base/`, you must regenerate the `assets_vfsdata.go` file using `go generate data/assets_generate.go`.
