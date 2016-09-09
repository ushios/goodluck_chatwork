package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ushios/goodluck_chatwork/command"
)

// DefaultCredentialFilePath .
var DefaultCredentialFilePath = "./.goodluck_chatwork"

// GlobalFlags .
var GlobalFlags = []cli.Flag{}

// Commands .
var Commands = []cli.Command{
	{
		Name:        "login",
		Description: "Create credential file",
		Action:      command.CmdLogin,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email",
				Usage: "Your registed email address",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "credential",
				Value: DefaultCredentialFilePath,
				Usage: "temporary file path",
			},
		},
	},
	{
		Name:        "log",
		Description: "List chat log",
		Action:      command.CmdLog,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email",
				Usage: "Your registed email address",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.Int64Flag{
				Name:  "room",
				Usage: "room or contact id",
			},
		},
	},
}

// CommandNotFound .
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
