package functions

import (
	"fmt"
	"strings"

	"github.com/Skrigbaum/Gobot/models"
)

//Card grabbed based on input
func Card(input string) string {
	var name string
	var skip bool
	var splitString = strings.Split(input, " ")

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
	//Default resonse if inproper command
	if name == "" {
		return "There seems to have been a problem with the command you entered, please try again."
	}

	var url = fmt.Sprintf("http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=" + name + "&type=card")
	name = ""
	skip = false
	return url
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
