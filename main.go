package main

import (
	"os"

	"github.com/benfb/vin/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vin"
	app.Usage = "the baseball command-line companion"
	app.Version = "0.2.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ben Bailey",
			Email: "bennettbailey@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "watch",
			Aliases:   []string{"w"},
			Usage:     "get notified when a blacked-out game is available",
			ArgsUsage: "[team] [phone]",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "interval, i",
					Value: 20,
					Usage: "how often to check if a game is over (in seconds)",
				},
			},
			Action: func(c *cli.Context) error {
				team := c.Args().Get(0)
				phone := c.Args().Get(1)
				if team != "" && phone != "" {
					commands.WatchClient(c.Uint64("interval"), team, phone)
				} else {
					return cli.NewExitError("Error! You must supply a team name and a phone number", 1)
				}
				return nil
			},
		},
		{
			Name:    "standings",
			Aliases: []string{"s"},
			Usage:   "get the current standings",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "aggregate, a",
					Usage: "get all standings in one table",
				},
			},
			Action: func(c *cli.Context) error {
				division := c.Args().Get(0)
				if division == "" && c.Bool("aggregate") {
					division = "agg"
				} else if division == "" {
					division = "all"
				}
				commands.StandingsCmd(division)
				return nil
			},
		},
		{
			Name:      "server",
			Aliases:   []string{"serve", "se"},
			Usage:     "run a vin server",
			ArgsUsage: "[address]",
			Action: func(c *cli.Context) error {
				commands.ServerCmd()
				return nil
			},
		},
		{
			Name:      "results",
			Aliases:   []string{"r"},
			Usage:     "get results for all the games from a particular day, formatted as mm/dd/yy",
			ArgsUsage: "date",
			UsageText: "vin results [-t team] date",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "team, t",
					Value: "all",
					Usage: "name of team to get box score for",
				},
				&cli.StringFlag{
					Name:  "except, e",
					Value: "",
					Usage: "name of team to exclude from results",
				},
			},
			Action: func(c *cli.Context) error {
				day := c.Args().Get(0)
				if day == "" {
					day = "today"
				}
				return commands.ResultsCmd(day, c.String("team"), c.String("except"))
			},
		},
	}

	app.Run(os.Args)
}
