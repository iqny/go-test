package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Version = "1.0.0"
	var IntValue int
	app.Usage = "this is a test project"
	app.Authors = []*cli.Author{
		{
			Name:  "lin",
			Email: "",
		},
	}
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Value:       8000,
			Usage:       "listening port",
			Destination: &IntValue,
		},
		&cli.StringFlag{
			Name:    "IP",
			Aliases: []string{"i", "I"},
			Value:   "127.0.0.1",
			Usage:   "listening IP",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
	}
	app.Commands = cli.Commands{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println(c.String("make"))
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "make",
					Aliases: []string{"m"},
					Value:   "127.0.0.1",
					Usage:   "listening IP",
				},
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("add 1")
				return nil
			},
		},
		{
			Name:    "queue",
			Aliases: []string{"q"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("queue")
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First(), c.String("make1"))
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "make1",
							Aliases: []string{"m"},
							Value:   "127.0.0.1",
							Usage:   "listening IP",
						},
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("new task template")
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello", c.Int("port"))
		fmt.Println(IntValue)
		fmt.Println(c.String("IP"))
		return nil
	}
	app.Run(os.Args)

	return
}
