package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/ushios/goodluck_chatwork/command"
)

// DefaultCredentialFilePath .
var DefaultCredentialFilePath = "./.goodluck_chatwork"

// GlobalFlags .
var GlobalFlags = []cli.Flag{}

// Commands .
var Commands = []cli.Command{
	{
		Name:        "list",
		Description: "Create credential file",
		Action:      command.CmdList,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email",
				Usage: "Your registed email address",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
		},
	},
	{
		Name:        "log",
		Description: "Save chat log to file",
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
			cli.IntFlag{
				Name:  "room",
				Usage: "room or contact id",
			},
		},
	},
	{
		Name:        "logall",
		Description: "Save all chat log to file",
		Action:      command.CmdLogAll,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email",
				Usage: "Your registed email address",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
		},
	},
}

// CommandNotFound .
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
