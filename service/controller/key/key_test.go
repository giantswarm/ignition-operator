package key

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	yaml "gopkg.in/yaml.v2"

	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/ignition-operator/service/controller/internal/unittest"
)

var update = flag.Bool("update", false, "update .golden CF template file")

// Test_Controller_Resource_TCNP_Template_Render tests tenant cluster
// CloudFormation template rendering. It is meant to be used as a tool to easily
// check resulting CF template and prevent from accidental CF template changes.
//
// It uses golden file as reference template and when changes to template are
// intentional, they can be updated by providing -update flag for go test.
//
//  go test ./service/controller/key -run Test_Controller_Key_Render -update
//
func Test_Controller_Key_Render(t *testing.T) {
	testCases := []struct {
		name string
		cc   interface{}
		path string
	}{
		{
			name: "case 0: file render",
			cc:   unittest.DefaultCC(),
			path: FilePath,
		},
		{
			name: "case 1: unit render",
			cc:   unittest.DefaultCC(),
			path: UnitPath,
		},
		{
			name: "case 2: master render",
			cc: func() controllercontext.Context {
				cc := unittest.DefaultCC()
				files, err := Render(cc, FilePath)
				if err != nil {
					panic(err)
				}
				units, err := Render(cc, UnitPath)
				if err != nil {
					panic(err)
				}
				cc.Status.Files = files
				cc.Status.Units = units
				return cc
			}(),
			path: MasterTemplatePath,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f, err := Render(tc.cc, tc.path)
			if err != nil {
				t.Fatal(err)
			}

			yamlFiles, err := yaml.Marshal(f)
			if err != nil {
				t.Fatal(err)
			}

			p := filepath.Join("testdata", unittest.NormalizeFileName(tc.name)+".golden")

			if *update {
				err := ioutil.WriteFile(p, yamlFiles, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := ioutil.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(yamlFiles, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(yamlFiles)))
			}
		})
	}
}
