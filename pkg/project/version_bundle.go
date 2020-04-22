package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "ignition-operator",
				Description: "Use Release.Revision in annotation for Helm 3 compatibility.",
				Kind:        versionbundle.KindChanged,
				URLs: []string{
					"https://github.com/giantswarm/ignition-operator/pull/68",
				},
			},
		},
		Components: []versionbundle.Component{},
		Name:       "ignition-operator",
		Version:    Version(),
	}
}
