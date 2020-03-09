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

	"github.com/giantswarm/ignition-operator/pkg/label"
	"github.com/giantswarm/ignition-operator/template/asset"
)

const (
	FilePath = "/files"
	UnitPath = "/units"

	MasterTemplatePath = "/ignition/master_template.yaml"
	WorkerTemplatePath = "/ignition/worker_template.yaml"
)

func OperatorVersion(getter LabelsGetter) string {
	return getter.GetLabels()[label.OperatorVersion]
}

func Render(values interface{}, filesdir string, b64 bool) (map[string]string, error) {
	files := make(map[string]string)

	walkFunction := func(path string, f os.FileInfo, rs io.ReadSeeker, err error) error {
		if err != nil {
			return microerror.Mask(err)
		}

		if !f.Mode().IsRegular() {
			return nil
		}

		file, err := vfsutil.ReadFile(asset.Assets, path)
		if err != nil {
			return microerror.Mask(err)
		}

		tmpl, err := template.New(path).Funcs(sprig.TxtFuncMap()).Parse(string(file))
		if err != nil {
			return microerror.Mask(err)
		}
		var data bytes.Buffer
		err = tmpl.Execute(&data, values)
		if err != nil {
			return microerror.Mask(err)
		}

		relativePath, err := filepath.Rel(filesdir, path)
		if err != nil {
			return microerror.Mask(err)
		}
		if b64 {
			files[relativePath] = base64.StdEncoding.EncodeToString(data.Bytes())
		} else {
			files[relativePath] = data.String()
		}

		return nil
	}

	err := vfsutil.WalkFiles(asset.Assets, filesdir, walkFunction)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return files, nil
}

func StatusSecretName(clusterID string, ignitionName string) string {
	return fmt.Sprintf("ignition-%s-%s", clusterID, ignitionName)
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
