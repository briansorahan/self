package si

import (
	"text/template"
)

func (t Tmpls) makefile() *template.Template {
	return template.Must(template.New("makefile").Parse(`APPNAME              = {{.Name}}
`))
}
