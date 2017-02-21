package functions

import (
	"fmt"
	"strings"

	"github.com/Skrigbaum/Gobot/models"
)

var (
	err      error
	adj      string
	race     string
	class    string
	location string
	quirk    string
	response string
)

//Fantasy checks secondary flag and redirect to proper function
func Fantasy(msg string) string {
	var splitString = strings.Split(msg, " ")
	switch flag := splitString[1]; flag {
	case "-p":
		response = Land()
	case "-c":
		response = Char()
	case "-q":
		response = Quirk()
	default:
		response = "Error reading flag. Try asking !gobot for help."
	}
	return response
}

//Name returns name type based on flag
func Name(msg []string) string {
	var name string
	var family string
	var nameType string

	if len(msg) == 1 {
		msg = append(msg, "-e")
	}
	switch flag := msg[1]; flag {
	case "-m":
		nameType = "male"
	case "-f":
		nameType = "female"
	case "-l":
		nameType = "family"
	case "-e":
		return "There seems to have been a problem finding you a name."
	}
	//Get name based on flag
	_ = models.DB.QueryRow("SELECT name from magic.names where type = '" + nameType + "' ORDER BY RAND() LIMIT 1").Scan(&name)

	//If flag is not family name then attatch last name to first name
	if nameType != "family" {
		_ = models.DB.QueryRow("SELECT name from magic.names where type = '" + nameType + "' ORDER BY RAND() LIMIT 1").Scan(&family)
		name = name + " " + family
	}
	return name

}

//Quirk genearetes random character quirks
func Quirk() string {
	var quirkErr = models.DB.QueryRow("SELECT QUIRK from magic.character where quirk != '' ORDER BY RAND() LIMIT 1").Scan(&quirk)
	if quirkErr != nil {
		return "There seems to have been a problem finding you a quirk."
	}
	return quirk
}

//Land generates a random land name
func Land() string {
	var landErr = models.DB.QueryRow("SELECT NAME FROM CARDS WHERE TYPE = 'LAND' ORDER BY RAND() LIMIT 1").Scan(&location)
	if landErr != nil {
		return "There seems to have been a problem finding you a place name."
	}
	return location
}

//Char is used to generate potential character ideas
func Char() string {
	//Queries need to be done like this due to bug in Mysql Driver for GO not supporting Stored Procedures
	_ = models.DB.QueryRow("SELECT adjective FROM magic.character where adjective != '' ORDER BY RAND() LIMIT 1;").Scan(&adj)
	_ = models.DB.QueryRow("SELECT race FROM magic.character where race != '' ORDER BY RAND() LIMIT 1;").Scan(&race)
	_ = models.DB.QueryRow("SELECT class FROM magic.character where class != '' ORDER BY RAND() LIMIT 1;").Scan(&class)
	_ = models.DB.QueryRow("SELECT location FROM magic.character where location != '' ORDER BY RAND() LIMIT 1;").Scan(&location)
	_ = models.DB.QueryRow("SELECT quirk FROM magic.character where quirk != '' ORDER BY RAND() LIMIT 1;").Scan(&quirk)
	var chars = fmt.Sprintf("Your character is a " + adj + " " + race + ". They are a " + class + " from the " + location + ", that " + quirk)
	chars = strings.ToLower(chars)
	return chars
}
