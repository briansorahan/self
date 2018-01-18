package si

import (
	"text/template"
)

func (t Tmpls) readme() *template.Template {
	return template.Must(template.New("readme").Parse(`# {{.Name}}

TODO: description of {{.Name}}

## Usage

### Start

` + "```" + `
make start
` + "```" + `

### Test

` + "```" + `
make test
` + "```" + `

### Push Docker Image

` + "```" + `
make push
` + "```" + `
`))
}
