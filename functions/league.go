package functions

import (
	"strconv"
	"strings"

	"github.com/TrevorSStone/goriot"
)

//League fetches RIOT games API data on current games
func League(input string) []string {
	var response []string
	var splitString = strings.Split(input, " ")

	//Default Card call
	if len(splitString) == 2 {
		//Assigns FeaturedGame to games variable
		normalizedName := goriot.NormalizeSummonerName(splitString[1])
		summoner, x := goriot.SummonerByName("NA", normalizedName[0])
		if x != nil {
			response = append(response, "Error Retrieving account")
		}
		var sumName = summoner[normalizedName[0]].Name
		response = append(response, sumName)

		matchHistory, matchErr := goriot.RecentGameBySummoner("na", summoner[normalizedName[0]].ID)
		if matchErr != nil {
			panic(matchErr.Error())
		}

		for _, game := range matchHistory {

			response = append(response, "Game ID: "+strconv.FormatInt(game.GameID, 10))
			response = append(response, "Playing: "+ChampName(game.ChampionID))
			response = append(response, "Score: "+strconv.Itoa(game.Statistics.ChampionsKilled)+"/"+strconv.Itoa(game.Statistics.NumDeaths)+"/"+strconv.Itoa(game.Statistics.Assists))
			if game.Statistics.NexusKilled == true {
				response = append(response, "You Lost \n")
			} else {
				response = append(response, "You Won \n")
			}
		}
	}
	return response
}

func Random() string {
	var char string
	for k := range Champions {
		return k
	}
	return char
}
