package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Skrigbaum/Gobot/models"
)

//Sets highest level
type Sets map[string]SetArray

//SetArray atruct for Card Set Arrays
type SetArray struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Cards []Cards `json:"cards"`
}

//Cards info
type Cards struct {
	ID           string   `json:"id"`
	ImageName    string   `json:"imageName"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Subtypes     []string `json:"subtypes"`
	Multiverseid int      `json:"multiverseid"`
}

//LoadCards is used to Load cads from the AllSets.json file into the mysql DB
func LoadCards() {
	//Read in file
	value, err := ioutil.ReadFile("Allsets.json")
	if err != nil {
		println(err)
	}
	//Unmarshal JSON into structs
	var sets Sets
	seterr := json.Unmarshal(value, &sets)
	if seterr != nil {
		print(err)
	}

	//Clean DB prior to run
	models.DB.Exec("Truncate TABLE CARDS;")

	//DB connection established

	//Iterate through sets structs and load requried info into DB
	for _, x := range sets {
		for _, y := range x.Cards {
			fmt.Println("Data being entered: " + y.ID + " " + strconv.Itoa(y.Multiverseid) + " " + y.ImageName + " " + y.Name + " " + y.Rarity + " " + y.Type + " ")
			if len(y.Subtypes) < 3 {
				y.Subtypes = append(y.Subtypes, "")
				y.Subtypes = append(y.Subtypes, "")
				y.Subtypes = append(y.Subtypes, "")
				fmt.Println(len(y.Subtypes))
			}
			result, queryerr := models.DB.Exec("INSERT INTO CARDS(ID, ImageName, MultiverseId, Name, Type, Rarity, SetCode, subtype, subtype2, subtype3) VALUES (?,?,?,?,?,?,?,?,?,?);",
				y.ID, y.ImageName, y.Multiverseid, y.Name, y.Type, y.Rarity, x.Code, y.Subtypes[0], y.Subtypes[1], y.Subtypes[2])
			if queryerr != nil {
				panic(queryerr)
			}
			println(result)
		}
	}

}
