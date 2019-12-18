package templatefiles

import (
	"bytes"
	"context"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/shurcooL/httpfs/vfsutil"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/ignition-operator/data"
	"github.com/giantswarm/ignition-operator/service/controller/controllercontext"
)

const (
	filesdir = "/files"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	cc.Status.Files, err = renderFiles(cc.Spec)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func renderFiles(spec controllercontext.ContextSpec) (map[string]string, error) {
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
