package functions

import (
	"strings"

	"github.com/TrevorSStone/goriot"
)

//League fetches RIOT games API data on current games
func League(input string) []string {
	var response []string
	var players []goriot.Participant
	var bans1 []int
	var bans2 []int
	var splitString = strings.Split(input, " ")

	//Default Card call
	if len(splitString) == 1 {

		//Assigns FeaturedGame to games variable
		var games, err = goriot.FeaturedGames("NA")
		if err != nil {
			panic(err.Error())
		}
		//Iterates through games to find players
		for _, p := range games[0].Participants {
			players = append(players, p)
		}
		//Goes and finds bans for the games and assigns to either team slice
		for _, b := range games[0].BannedChampions {
			if b.TeamID == 1 {
				bans1 = append(bans1, b.ChampionID)
			} else {
				bans2 = append(bans2, b.ChampionID)
			}
		}

		//begin creating response
		response = append(response, "Match ID: "+string(games[0].GameID))
		response = append(response, "Current Game Duration: "+string(games[0].GameLength))
		response = append(response, "Team 1 Banned: ")
		for _, ban1 := range bans1 {
			response = append(response, string(ban1))
		}
		response = append(response, "Team 2 Banned: ")
		for _, ban2 := range bans2 {
			response = append(response, string(ban2))
		}
		for _, player := range players {
			response = append(response, "Summoner: "+player.SummonerName)
			response = append(response, "On: "+goriot.ChampionByID("NA", player.ChampionID))
			response = append(response, ban2)

		}
	}
}
