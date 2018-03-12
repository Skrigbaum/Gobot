package functions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Skrigbaum/Gobot/models"
)

//Card grabbed based on input
func Card(input string) string {
	var name string
	var setting string
	var antagonist string
	var helper string
	var problem string
	var solution string
	var skip bool
	var splitString = strings.Fields(input)

	//Default Card call
	if len(splitString) == 1 {
		err = models.DB.QueryRow("SELECT MULTIVERSEID FROM CARDS ORDER BY RAND() LIMIT 1").Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		skip = true
	}
	//Type check
	if !skip && splitString[1] == "type" && len(splitString) > 2 {
		var typeErr = models.DB.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE TYPE LIKE '" + splitString[2] + "%' ORDER BY RAND() LIMIT 1").Scan(&name)
		if typeErr != nil {
			return "There seems to have been a problem with the type you entered, please try again."
		}
	}
	//subtype check
	if !skip && splitString[1] == "subtype" && len(splitString) > 2 {
		var typeErr = models.DB.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE SUBTYPE = '" + splitString[2] + "' OR SUBTYPE2 = '" + splitString[2] + "' OR SUBTYPE3 = '" + splitString[2] + "' ORDER BY RAND() LIMIT 1").Scan(&name)
		if typeErr != nil {
			return "There seems to have been a problem with the type you entered, please try again."
		}
	}
	//Set check
	if !skip && splitString[1] == "set" && len(splitString) > 2 {
		var typeErr = models.DB.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE SETCODE = '" + splitString[2] + "' ORDER BY RAND() LIMIT 1").Scan(&name)
		if typeErr != nil {
			return "There seems to have been a problem with the type you entered, please try again."
		}
	}

	//Name check
	if !skip && splitString[1] == "rarity" && len(splitString) > 2 {
		var nameErr = models.DB.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE rarity like '" + splitString[2] + "%' ORDER BY RAND() LIMIT 1").Scan(&name)
		if nameErr != nil {
			return "There seems to have been a problem with the rarity you entered, please try again."
		}
	}

	//Adventure check
	if !skip && splitString[1] == "adv" {
		if len(splitString) > 2 {
			_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Artifact%' or type like '%Sorcery%' AND setcode = '" + splitString[2] + "' order by RAND() limit 1;").Scan(&solution)
			_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%enchantment%' or type like '%conspiracy%' or type like '%Phenomenon%' or type like '%Planeswalker%' or type like '%Vanguard%' AND setcode = '" + splitString[2] + "' order by RAND() limit 1;").Scan(&problem)
			_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Land%' AND setcode = '" + splitString[2] + "' order by RAND() limit 1;").Scan(&setting)
			_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%PlanesWalker%' AND setcode = '" + splitString[2] + "' order by RAND() limit 1;").Scan(&helper)
			_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%PlanesWalker%' AND setcode = '" + splitString[2] + "' order by RAND() limit 1;").Scan(&antagonist)

			test1 := APICall(solution)
			test2 := APICall(problem)
			test3 := APICall(setting)
			test4 := APICall(helper)
			test5 := APICall(antagonist)

			var response = fmt.Sprintf("Problem: " + test2 + "\n" + "Setting: " + test3 + "\n" + "Helper: " + test4 + "\n" + "Possible Solution: " + test1 + "\n" + "Antagonist: " + test5)
			return response
		}
		_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Artifact%' or type like '%Sorcery%' order by RAND() limit 1;").Scan(&solution)
		_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%enchantment%' or type like '%conspiracy%' or type like '%Phenomenon%' or type like '%Planeswalker%' or type like '%Vanguard%' order by RAND() limit 1;").Scan(&problem)
		_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Land%' order by RAND() limit 1;").Scan(&setting)
		_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%PlanesWalker%' order by RAND() limit 1;").Scan(&helper)
		_ = models.DB.QueryRow("Select MULTIVERSEID from MTG.cards where type like '%Creature%' or type like '%PlanesWalker%' order by RAND() limit 1;").Scan(&antagonist)

		test1 := APICall(solution)
		test2 := APICall(problem)
		test3 := APICall(setting)
		test4 := APICall(helper)
		test5 := APICall(antagonist)

		var response = fmt.Sprintf("Problem: " + test2 + "\n" + "Setting: " + test3 + "\n" + "Helper: " + test4 + "\n" + "Possible Solution: " + test1 + "\n" + "Antagonist: " + test5)
		return response

	}
	//Default resonse if inproper command
	if name == "" {
		return "There seems to have been a problem with the command you entered, please try again."
	}

	respURL := APICall(name)
	name = ""
	skip = false
	return respURL
}

//Place Generates Place Name
func Place() string {
	var placeName string
	err = models.DB.QueryRow("SELECT NAME FROM CARDS WHERE TYPE = 'LAND' ORDER BY RAND() LIMIT 1").Scan(&placeName)
	if err != nil {
		panic(err.Error())
	}
	return placeName
}

//APICall Exports API call for future extensions
func APICall(name string) string {
	var url2 = fmt.Sprintf("https://api.scryfall.com/cards/multiverse/" + name)
	req, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		panic(err.Error())
	}
	test := req.URL.Query()
	test.Add(":id", "url")
	test.Add("format", "image")
	test.Add("version", "large")
	req.URL.RawQuery = test.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Processing...")
		data := resp.Request.URL.String()
		fmt.Println(data)
		return string(data)
	}

}
