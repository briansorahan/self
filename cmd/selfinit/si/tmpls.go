package si

import (
	"os"
	"text/template"

	"github.com/pkg/errors"
)

// Tmpls renders templates.
type Tmpls struct {
	m map[string]*template.Template
}

// NewTmpls creates templates.
func NewTmpls(name string) (Tmpls, error) {
	t := Tmpls{}

	t.m = map[string]*template.Template{
		"README.md":          t.readme(),
		"Dockerfile":         t.dockerfile(),
		"Dockerfile.release": t.release(),
		"Makefile":           t.makefile(),
		".gitignore":         t.gitignore(),
		"app.go":             t.app(),
		"config.go":          t.config(),
		"config_test.go":     t.configtest(),
		"main.go":            t.main(),
	}
	return t, nil
}

// WriteAll writes all the templates to disk.
func (t Tmpls) WriteAll(data Data) error {
	var files []*os.File

	defer func() {
		for _, f := range files {
			_ = f.Close() // Best effort.
		}
	}()
	for filename, tmpl := range t.m {
		f, err := os.Create(filename)
		if err != nil {
			return errors.Wrap(err, "creating file")
		}
		files = append(files, f)

		if err := tmpl.Execute(f, data); err != nil {
			return errors.Wrap(err, "executing template")
		}
	}
	return nil
}
