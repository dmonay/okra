package main

import (
	// "fmt"

	"github.com/codegangsta/cli"
	"github.com/dmonay/okra/handlers"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "Do-Work"
	app.Usage = "Personal JIRA board"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", ""},
	}

	app.Commands = handlers.Commands

	app.Run(os.Args)

}
