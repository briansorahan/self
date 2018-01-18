package si

import (
	"text/template"
)

func (t Tmpls) gitignore() *template.Template {
	return template.Must(template.New("gitignore").Parse(`.assets.tar
.build
.push
.release
`))
}
