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
		name   string
		base64 bool
		cc     interface{}
		path   string
		print  bool
	}{
		{
			name:   "case 0: file render",
			base64: false,
			cc:     unittest.DefaultCC(),
			path:   FilePath,
		},
		{
			name:   "case 1: unit render",
			base64: false,
			cc:     unittest.DefaultCC(),
			path:   UnitPath,
		},
		{
			name:   "case 2: master render",
			base64: false,
			cc: func() controllercontext.Context {
				cc := unittest.DefaultCC()
				files, err := Render(cc, FilePath, true)
				if err != nil {
					panic(err)
				}
				units, err := Render(cc, UnitPath, false)
				if err != nil {
					panic(err)
				}
				cc.Status.Files = files
				cc.Status.Units = units
				return cc
			}(),
			path:  MasterTemplatePath,
			print: true,
		},
		{
			name:   "case 3: worker render",
			base64: false,
			cc: func() controllercontext.Context {
				cc := unittest.DefaultCC()
				files, err := Render(cc, FilePath, true)
				if err != nil {
					panic(err)
				}
				units, err := Render(cc, UnitPath, false)
				if err != nil {
					panic(err)
				}
				cc.Status.Files = files
				cc.Status.Units = units
				return cc
			}(),
			path:  WorkerTemplatePath,
			print: true,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f, err := Render(tc.cc, tc.path, tc.base64)
			var yamlFiles []byte

			if err != nil {
				t.Fatal(err)
			}
			if tc.print {
				yamlFiles = []byte(f["."])
			} else {
				yamlFiles, err = yaml.Marshal(f)
				if err != nil {
					t.Fatal(err)
				}
			}

			p := filepath.Join("testdata", unittest.NormalizeFileName(tc.name)+".golden")

			if *update {
				err := ioutil.WriteFile(p, yamlFiles, 0644) // #nosec
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
