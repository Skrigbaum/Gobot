package functions

import (
	"fmt"

	"github.com/Skrigbaum/Gobot/models"
)

var (
	err error
)

//Char is used to generate potential character ideas
func Char() string {
	var adj string
	var race string
	var class string
	var location string
	var quirk string

	//Queries need to be done like this due to bug in Mysql Driver for GO not supporting Stored Procedures
	_ = models.DB.QueryRow("SELECT adjective FROM magic.character where adjective != '' ORDER BY RAND() LIMIT 1;").Scan(&adj)
	_ = models.DB.QueryRow("SELECT race FROM magic.character where race != '' ORDER BY RAND() LIMIT 1;").Scan(&race)
	_ = models.DB.QueryRow("SELECT class FROM magic.character where class != '' ORDER BY RAND() LIMIT 1;").Scan(&class)
	_ = models.DB.QueryRow("SELECT location FROM magic.character where location != '' ORDER BY RAND() LIMIT 1;").Scan(&location)
	_ = models.DB.QueryRow("SELECT quirk FROM magic.character where quirk != '' ORDER BY RAND() LIMIT 1;").Scan(&quirk)

	var chars = fmt.Sprintf("Your character is a " + adj + " " + race + ". They are a " + class + " from the " + location + ", that " + quirk)
	return chars

}
