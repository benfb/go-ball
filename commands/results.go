package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/benfb/vin/api"
	"github.com/benfb/vin/util"
	"github.com/urfave/cli/v2"
)

// ResultsCmd is the command run by `vin results`
func ResultsCmd(date, team, without string) error {
	if !util.ContainsString(api.Teams, strings.Title(team)) && team != "all" {
		return cli.NewExitError("Error! \""+team+"\" is not a valid team.", 1)
	}
	timeFmtStr := "1/_2/06"
	if date == "today" {
		date = time.Now().Format(timeFmtStr)
	}
	parsedTime, timeErr := time.Parse(timeFmtStr, date)
	if timeErr != nil {
		return cli.NewExitError("Error! \""+date+"\" is not a valid date.", 1)
	}
	list := api.FetchGames(parsedTime)
	for _, g := range list {
		if !g.FindTeam(strings.Title(without)) && (g.FindTeam(strings.Title(team)) || team == "all") {
			api.PrintBoxScoreTable(g, api.FetchLineScore(strconv.Itoa(g.GamePk)))
		}
	}

	return nil
}
