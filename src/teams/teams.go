package teams

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Team struct {
	Name     string
	Total    string
	Spread   string
	ML       string
	Opponent *Team
}

func GetTeamInfoForSport(sport string) []Team {

	c := colly.NewCollector()
	teamNames := make([]string, 0)
	lines := make([]string, 0)

	//scrape all of the line info
	c.OnHTML("div.market", func(e *colly.HTMLElement) {
		lines = append(lines, e.Text)
	})

	//get team names
	c.OnHTML(".team-title", func(e *colly.HTMLElement) {
		teamNames = append(teamNames, e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	teams := make([]Team, 0)

	sportToURL := map[string]string{"NFL": "https://www.sportsbook.ag/sbk/sportsbook4/nfl-betting/nfl-lines-nfl-game-lines.sbk",
		"NBA":   "https://www.sportsbook.ag/sbk/sportsbook4/nba-betting/nba-game-lines-nba-game-lines.sbk",
		"NHL":   "https://www.sportsbook.ag/sbk/sportsbook4/nhl-betting/nhl-game-lines.sbk",
		"NCAAB": "https://www.sportsbook.ag/sbk/sportsbook4/ncaab-betting/game-lines.sbk"}

	c.Visit(sportToURL[sport])

	var first bool = true
	lineIdx := 0 //starts at total (O/U), then spread, ML
	for i, team := range teamNames {
		var newTeam Team
		newTeam.Name = team
		newTeam.Total = lines[lineIdx]
		newTeam.Spread = lines[lineIdx+1]
		newTeam.ML = lines[lineIdx+2]
		lineIdx += 3
		if first == false {
			newTeam.Opponent = &(teams[i-1])
			teams[i-1].Opponent = &newTeam
			first = true
		} else {
			first = false
		}
		teams = append(teams, newTeam)
	}

	return teams
}

func PrintInfoForTeams(teams []Team) string {

	strBuild := ""

	//sportsbook always lists away team first
	i := 0
	for i < len(teams) {
		toPrint := fmt.Sprintf("%s @ %s\n", teams[i].Name, teams[i].Opponent.Name)
		strBuild += toPrint
		dashes := len(toPrint) - 1 //don't want the newline char
		for dashes > 0 {
			strBuild += fmt.Sprintf("-")
			dashes--
		}
		strBuild += "\n"

		strBuild += fmt.Sprintf("%s: %s, %s, %s\n", teams[i].Name, teams[i].ML, teams[i].Spread, teams[i].Total)
		strBuild += fmt.Sprintf("%s: %s, %s, %s\n", teams[i+1].Name, teams[i+1].ML, teams[i+1].Spread, teams[i+1].Total)

		//line seperation
		strBuild += "\n"
		strBuild += "\n"
		i += 2
	}

	return strBuild
}

func FilterTeamsForFavorites(teams []Team) []Team {

	favoriteTeamsSet := map[string]int{"Dallas Cowboys": 1, "Dallas Mavericks": 1, "Dallas Stars": 1, "Illinois": 1, "Oklahoma": 1,
		"Chicago Bulls": 1, "Chicago Bears": 1, "Chicago Blackhawks": 1}

	var favTeams []Team

	for _, team := range teams {
		_, okTeam := favoriteTeamsSet[team.Name]
		_, okTeamOpp := favoriteTeamsSet[team.Opponent.Name]
		if okTeam || okTeamOpp {
			favTeams = append(favTeams, team)
		}
	}

	return favTeams

}
