package hanlonlib

import (
	"github.com/codegangsta/cli"
	"fmt"
)

func NodeCommands() cli.Command {
	command := cli.Command{
		Name: "node",
		Usage: "Show available nodes",
		Action: func(c *cli.Context) { fmt.Println("print images") },
		Subcommands: []cli.Command{
			{
				Name: "update",
				Usage: "Remove an image",
				Action: func(c *cli.Context) {
					println("removed image")
				},
			},
		},
	}
	return command
}

