package si

import (
	"text/template"
)

func (t Tmpls) release() *template.Template {
	return template.Must(template.New("dockerfile.release").Parse(`FROM       scratch
ADD	   .assets.tar /
ENTRYPOINT ["{{.BuildOutput}}"]
`))
}
