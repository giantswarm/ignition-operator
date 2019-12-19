package key

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path/filepath"

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

func OperatorVersion(getter LabelsGetter) string {
	return getter.GetLabels()[label.OperatorVersion]
}

func Render(values interface{}, filesdir string) (map[string]string, error) {
	files := make(map[string]string)

	err := vfsutil.WalkFiles(data.Assets, filesdir, func(path string, f os.FileInfo, rs io.ReadSeeker, err error) error {
		if !f.Mode().IsRegular() {
			return nil
		}

		file, err := vfsutil.ReadFile(data.Assets, path)
		if err != nil {
			return microerror.Mask(err)
		}

		tmpl, err := template.New(path).Funcs(sprig.FuncMap()).Parse(string(file))
		if err != nil {
			return microerror.Mask(err)
		}
		var data bytes.Buffer
		tmpl.Execute(&data, values)

		relativePath, err := filepath.Rel(filesdir, path)
		if err != nil {
			return microerror.Mask(err)
		}
		files[relativePath] = string(data.Bytes())

		return nil
	})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return files, nil
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
