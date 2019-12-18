package templatefiles

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
//  go test ./service/controller/resource/templatefiles -run Test_Controller_Resource_Templatefiles_Renderfiles -update
//
func Test_Controller_Resource_Templatefiles_Renderfiles(t *testing.T) {
	testCases := []struct {
		name   string
		ccSpec controllercontext.ContextSpec
	}{
		{
			name: "case 0: basic test",
			ccSpec: controllercontext.ContextSpec{
				BaseDomain: "someshit",
				Etcd: controllercontext.ContextSpecEtcd{
					Domain: "abc",
				},
				Kubernetes: controllercontext.ContextSpecKubernetes{
					DNS: controllercontext.ContextSpecKubernetesDNS{
						IP: "k8sdnsIP",
					},
					Domain: "k8sdomain",
				},
				SSO: controllercontext.ContextSpecSSO{
					PublicKey: "some secret",
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f, err := renderFiles(tc.ccSpec)
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
