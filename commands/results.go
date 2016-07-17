package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/benfb/vin/api"
	"github.com/benfb/vin/util"
	"github.com/urfave/cli"
)

// ResultsCmd is the command run by `vin results`
func ResultsCmd(date, team string) error {
	go util.Spinner()

	log.Println(strings.Title(team))
	if !util.ContainsString(api.Teams, strings.Title(team)) {
		return cli.NewExitError("Error! \""+team+"\" is not a valid team.", 1)
	}
	timeFmtStr := "1/_2/06"
	if date == "today" {
		date = time.Now().Format(timeFmtStr)
	}
	parsedTime, timeErr := time.Parse(timeFmtStr, date)
	if timeErr != nil {
		log.Fatalln("That is not a valid date!")
	}
	list := api.FetchGames(parsedTime)
	for _, g := range list {
		if g.FindTeam(strings.Title(team)) || team == "all" {
			g.PrintBoxScoreTable()
			fmt.Println("Inning: " + util.FormatInning(g.Inning, g.IsTop, g.Status))
		}
	}

	return nil
}
