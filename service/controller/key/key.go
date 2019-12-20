package key

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/shurcooL/httpfs/vfsutil"

	"github.com/giantswarm/ignition-operator/data"
	"github.com/giantswarm/ignition-operator/pkg/label"
)

const (
	FilePath = "/files"
	UnitPath = "/units"

	MasterTemplatePath = "master_template.yaml"
	WorkerTemplatePath = "worker_template.yaml"
)

const (
	DefaultNamespace = "giantswarm"
)

func OperatorVersion(getter LabelsGetter) string {
	return getter.GetLabels()[label.OperatorVersion]
}

func Render(values interface{}, filesdir string, b64 bool) (map[string]string, error) {
	files := make(map[string]string)

	err := vfsutil.WalkFiles(data.Assets, filesdir, func(path string, f os.FileInfo, rs io.ReadSeeker, err error) error {
		if !f.Mode().IsRegular() {
			return nil
		}

		file, err := vfsutil.ReadFile(data.Assets, path)
		if err != nil {
			return microerror.Mask(err)
		}

		tmpl, err := template.New(path).Funcs(sprig.TxtFuncMap()).Parse(string(file))
		if err != nil {
			return microerror.Mask(err)
		}
		var data bytes.Buffer
		tmpl.Execute(&data, values)

		relativePath, err := filepath.Rel(filesdir, path)
		if err != nil {
			return microerror.Mask(err)
		}
		if b64 {
			files[relativePath] = base64.StdEncoding.EncodeToString(data.Bytes())
		} else {
			files[relativePath] = string(data.Bytes())
		}

		return nil
	})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return files, nil
}

func StatusConfigMapName(clusterID string) string {
	return fmt.Sprintf("ignition-operator-%s", clusterID)
}

func ToIgnition(v interface{}) (v1alpha1.Ignition, error) {
	if v == nil {
		return v1alpha1.Ignition{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.Ignition{}, v)
	}

	p, ok := v.(*v1alpha1.Ignition)
	if !ok {
		return v1alpha1.Ignition{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.Ignition{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}
