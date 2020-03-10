//go:generate go run -tags=dev generate.go
// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/giantswarm/ignition-operator/pkg/asset"
)

func main() {
	err := vfsgen.Generate(asset.Assets, vfsgen.Options{
		BuildTags:    "!dev",
		Filename:     "vfsstatic.go",
		PackageName:  "asset",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
