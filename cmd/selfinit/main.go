package main

import (
	"context"
	"flag"

	"github.com/briansorahan/self/cmd/selfinit/si"
)

func main() {
	app, err := si.NewApp("selfinit", flag.ContinueOnError)
	si.Die(err) // Only dies if err != nil
	si.Die(app.Run(context.Background()))
}
