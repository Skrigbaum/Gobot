package functions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/justinian/dice"
)

//Rolling is to simulate rolling D&D dice
func Rolling(input string) string {
	var splitString = strings.Split(input, " ")
	//if len(splitString) <= 1 {
	//	return "Please enter dice equation. ex: !roll 1d6+2. You input: " + splitString[0]
	//}
	result, _, err := dice.Roll(splitString[1])
	if err != nil {
		return "error"
	}
	re := regexp.MustCompile("\\[(.*?)\\]")
	var parsed = re.FindString(result.String())

	return "Results: " + parsed + " " + result.Description() + "\n" + "Total: " + strconv.Itoa(result.Int())
}
