package functions

import (
	"strconv"
	"strings"

	"github.com/TrevorSStone/goriot"
)

//League fetches RIOT games API data on current games
func League(input string) []string {
	var response []string
	var players []goriot.Participant
	var bans []string
	var splitString = strings.Split(input, " ")

	//Default Card call
	if len(splitString) == 1 {
		//Assigns FeaturedGame to games variable
		games, x := goriot.FeaturedGames("NA")
		if x != nil {
			panic(x.Error())
		}
		//Iterates through games to find players
		for _, p := range games[0].Participants {
			players = append(players, p)
		}
		//Goes and finds bans for the games and assigns to either team slice
		for _, b := range games[0].BannedChampions {
			name := ChampName(b.ChampionID)
			bans = append(bans, name)
		}

		//begin creating response
		response = append(response, "Current Game Duration: "+strconv.Itoa(games[0].GameLength))
		response = append(response, "Champions Banned: ")
		for _, ban1 := range bans {
			response = append(response, ban1)
		}
		for _, p := range players {
			response = append(response, "Summoner: "+p.SummonerName)
			response = append(response, "On: "+ChampName(p.ChampionID))
			response = append(response, ("Player Score: " + string(p.Stats.Kills) + "/" + string(p.Stats.Deaths) + "/" + string(p.Stats.Assists)))
		}
	}
	return response
}
