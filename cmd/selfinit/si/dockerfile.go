package si

import (
	"text/template"
)

func (t Tmpls) dockerfile() *template.Template {
	return template.Must(template.New("dockerfile").Parse(`FROM		{{.BuildImage}}
ARG		APPNAME
ENV		APP_PATH $GOPATH/src/$APPNAME
WORKDIR		$APP_PATH
ADD		. $APP_PATH
RUN		go get -d ./...
RUN             go test -race
RUN             CGO_ENABLED=0 go build {{.BuildFlags}} -o {{.BuildOutput}}
`))
}
