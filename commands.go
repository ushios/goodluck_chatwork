package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// GlobalFlags .
var GlobalFlags = []cli.Flag{}

// Commands .
var Commands = []cli.Command{}

// CommandNotFound .
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
