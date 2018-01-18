package main

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
