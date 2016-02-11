package main

import (
	"os"
	"github.com/codegangsta/cli"
	"hanloncli/hanlonlib"
)
func main()  {
	app := cli.NewApp()
	commands := []cli.Command{
		hanlonlib.ImageCommands(),
		hanlonlib.ModelCommands(),
		hanlonlib.NodeCommands(),
	}
	app.Commands = commands
	app.Run(os.Args)
}
