package si

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/pkg/errors"
)

// Flag defaults.
const (
	defaultBuildFlags = `-ldflags "-d -X main.Version=$VERSION -X main.BuildCommitID=$BUILD_GIT_COMMIT_ID -X main.BuildTime=$BUILD_TIME -X main.BuildURI=$BUILD_URL"`
	defaultBuildImage = "bsorahan/go-build-image"
)

// App defines the behavior of the application.
type App struct {
	Data

	flagset *flag.FlagSet
	pass    bool // If true the application does nothing.
	tmpls   Tmpls
}

// NewApp creates a new applictiona.
func NewApp(name string, fe flag.ErrorHandling) (*App, error) {
	const permsFlagName = "perms"
	// Parse flags.
	a := &App{
		flagset: flag.NewFlagSet(name, fe),
	}
	a.flagset.StringVar(&a.BuildFlags, "build-flags", defaultBuildFlags, "Build flags.")
	a.flagset.StringVar(&a.BuildImage, "build-image", defaultBuildImage, "Base image for build.")
	a.flagset.StringVar(&a.BuildOutput, "build-output", "/app", "Path to binary in build image.")
	a.flagset.StringVar(&a.Name, "name", "", "(required) Application name.")
	a.flagset.StringVar(&a.Org, "org", "", "(required) Github organization.")
	a.flagset.StringVar(&a.Repo, "repo", "bsorahan/"+a.Name, "Dockerhub repo name.")

	if err := a.flagset.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			a.pass = true
			return a, nil
		}
		return nil, errors.Wrap(err, "parsing flags")
	}
	if err := a.ValidateFlags(); err != nil {
		return nil, errors.Wrap(err, "invalid flag(s)")
	}
	// Create templates.
	tmpls, err := NewTmpls(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, "creating templates")
	}
	a.tmpls = tmpls

	return a, nil
}

// Run runs the application.
func (a *App) Run(ctx context.Context) error {
	if a.pass {
		return nil
	}
	return errors.Wrap(a.tmpls.WriteAll(a.Data), "writing templates")
}

// ValidateFlags validates the application's flags.
func (a *App) ValidateFlags() error {
	if len(a.Name) == 0 {
		return errors.New("-name is required")
	}
	if len(a.Org) == 0 {
		return errors.New("-org is required")
	}
	return nil
}

// Die kills the program.
func Die(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

// Data is used to render all the templates.
type Data struct {
	BuildFlags  string
	BuildImage  string
	BuildOutput string
	JobType     string
	Name        string
	Org         string
	Perms       uint
	Repo        string
}

func (t Tmpls) app() *template.Template {
	return template.Must(template.New("main").Parse(`package main

import (
	"context"
	"flag"
	"os"

	"github.com/pkg/errors"
)

// App defines the behavior of the application.
type App struct {
	Config

	pass bool
}

var flagErrorHandling = flag.ContinueOnError

// NewApp creates a new applictiona.
func NewApp(name string) (*App, error) {
	var (
		config  Config
		flagset = flag.NewFlagSet(name, flagErrorHandling)
	)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return &App{pass: true}, nil
		}
		return nil, errors.Wrap(err, "parsing flags")
	}
	return NewAppFrom(name, config)
}

// NewAppFrom creates a new app from the provided config.
func NewAppFrom(name string, config Config) (*App, error) {
	a := &App{
		Config: config,
	}
	if err := a.Config.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid flag(s)")
	}
	return a, nil
}

// Run runs the application.
func (a *App) Run(ctx context.Context) error {
	if a.pass {
		return nil
	}
	// TODO
	return nil
}
`))
}
