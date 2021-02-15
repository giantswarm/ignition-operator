module github.com/giantswarm/ignition-operator

go 1.13

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/giantswarm/apiextensions/v3 v3.17.0
	github.com/giantswarm/k8sclient/v5 v5.10.0
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/operatorkit/v4 v4.2.0
	github.com/giantswarm/versionbundle v0.2.0
	github.com/google/go-cmp v0.5.4
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/spf13/viper v1.7.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.15
	k8s.io/apimachinery v0.18.15
	k8s.io/client-go v0.18.15
)

replace gopkg.in/fsnotify.v1 v1.4.7 => gopkg.in/fsnotify/fsnotify.v1 v1.4.7

replace github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2

replace github.com/coreos/etcd v3.3.17+incompatible => github.com/coreos/etcd v3.3.25+incompatible

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible

replace github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible

// keep in sync with giantswarm/apiextensions
replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
