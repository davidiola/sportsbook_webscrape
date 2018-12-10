package main

import (
	"github.com/davidiola/sportsbook_webscrape/src/teams"
	"github.com/davidiola/sportsbook_webscrape/src/twilio"
)

func main() {

	var allTeams []teams.Team

	nflTeams := teams.GetTeamInfoForSport("NFL")
	allTeams = append(nflTeams, allTeams...)
	nbaTeams := teams.GetTeamInfoForSport("NBA")
	allTeams = append(nbaTeams, allTeams...)
	ncaabTeams := teams.GetTeamInfoForSport("NCAAB")
	allTeams = append(ncaabTeams, allTeams...)
	nhlTeams := teams.GetTeamInfoForSport("NHL")
	allTeams = append(nhlTeams, allTeams...)

	favoriteTeams := teams.FilterTeamsForFavorites(allTeams)

	messageBody := teams.PrintInfoForTeams(favoriteTeams)

	twilio.SendTextWithMessage(messageBody)

}
