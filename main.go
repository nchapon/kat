package main

import (
	"fmt"
	"kat/commands"
	"os"

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
			Name:    "setenv",
			Aliases: []string{"s"},
			Usage:   "Set up kong env",
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
