package main

import (
	"fmt"
	"os"

	"github.com/nchapon/kat/commands"
	"github.com/urfave/cli"
)

// main ...
func main() {

	app := cli.NewApp()
	app.Name = "KAT"
	app.Description = "Kat Admin Tools"

	app.Commands = []cli.Command{
		{
			Name:    "add-consumer",
			Aliases: []string{"ac"},
			Usage:   "Add a new CONSUMER",
			Action: func(c *cli.Context) error {
				fmt.Println("added consumer : ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "Deploy a new API",
			Action: func(c *cli.Context) error {
				fmt.Println("Deploying API", c.Args().First())
				return nil
			},
		},
		{
			Name:    "init",
			Aliases: []string{"s"},
			Usage:   "Init KAT project",
			Action: func(c *cli.Context) error {
				commands.Init()
				return nil
			},
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Unknown command  %q here.\n", command)
	}

	app.Run(os.Args)
}
