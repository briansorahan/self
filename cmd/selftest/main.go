package main

import (
	"context"
	"fmt"
	"os"
)

// AppName is the name of the application.
const AppName = "selftest"

func main() {
	app, err := NewApp(AppName)
	Die(err) // Only dies if err != nil
	Die(app.Run(context.Background()))
}

// Die kills the program.
func Die(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
