package hanlonlib

import (
	"github.com/codegangsta/cli"
	"fmt"
)

func ModelCommands() cli.Command {
	command := cli.Command{
		Name: "model",
		Usage: "Show available models",
		Action: func(c *cli.Context) { fmt.Println("print images") },
		Subcommands: []cli.Command {
			{
				Name: "remove",
				Usage: "Remove an image",
				Action: func(c *cli.Context) {
					println("removed image")
				},
			},
			{
				Name: "add",
				Usage: "Add a model",
				Action: func(c *cli.Context) {
					println("Add image")
				},
			},
			{
				Name: "update",
				Usage: "update a model",
				Action: func(c *cli.Context) {
					println("Add image")
				},
			},
		},
	}
	return command
}

