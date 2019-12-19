package key

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/ignition-operator/data"
	"github.com/giantswarm/ignition-operator/pkg/label"
	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
	"github.com/giantswarm/microerror"
	"github.com/shurcooL/httpfs/vfsutil"
)

const (
	FilePath = "/files"
	UnitPath = "/units"
)

func OperatorVersion(getter LabelsGetter) string {
	return getter.GetLabels()[label.OperatorVersion]
}

func Render(spec controllercontext.ContextSpec, filesdir string) (map[string]string, error) {
	files := make(map[string]string)

	err := vfsutil.WalkFiles(data.Assets, filesdir, func(path string, f os.FileInfo, rs io.ReadSeeker, err error) error {
		if f.Mode().IsRegular() {
			file, err := vfsutil.ReadFile(data.Assets, path)
			if err != nil {
				return microerror.Mask(err)
			}

			tmpl, err := template.New(path).Parse(string(file))
			if err != nil {
				return microerror.Maskf(err, "failed to parse file %#q", path)
			}
			var data bytes.Buffer
			tmpl.Execute(&data, spec)

			relativePath, err := filepath.Rel(filesdir, path)
			if err != nil {
				return microerror.Mask(err)
			}
			files[relativePath] = string(data.Bytes())
		}
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
