package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	buffer = make([][]byte, 0)
	Token  string
	BotID  string
	db     *sql.DB
	err    error
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	if err != nil {
		fmt.Println("Error loading sound: ", err)
		fmt.Println("Please copy $GOPATH/src/github.com/bwmarrin/examples/airhorn/airhorn.dca to this directory.")
		return
	}
	// Create a new Discord session using the provided bot token.
	dg, errCon := discordgo.New("Bot " + Token)
	if errCon != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	var msg = strings.ToLower(m.Content)
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if msg == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if msg == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	//if the message is !gobot return a helpful message
	if msg == "!gobot" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi I'm Gobot. Currently typing '!card returns a random magic card', '!card type *X* returns a card of that type', '!card rarity *X* returns a card of that rarity'. \n!char returns a random D&D character.")
	}

	if strings.Contains(msg, "!card") {
		cardName := Mtg(msg)
		if err != nil {
			fmt.Println("Error retriving card,", err)
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, cardName)

	}

	if strings.Contains(msg, "!char") {
		character := Char()
		if err != nil {
			fmt.Println("Error retriving character,", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Stop breaking stuff.")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, character)

	}
}

//Mtg Grabs card from DB
func Mtg(input string) string {
	var name string
	var skip bool
	var splitString = strings.Split(input, " ")

	db, err = sql.Open("mysql", "root:1Merdenomz1@/magic")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Default Card call
	if len(splitString) == 1 {
		err = db.QueryRow("SELECT MULTIVERSEID FROM CARDS ORDER BY RAND() LIMIT 1").Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		skip = true
	}
	//Type check
	if !skip && splitString[1] == "type" && len(splitString) > 2 {
		var typeErr = db.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE TYPE = '" + splitString[2] + "' ORDER BY RAND() LIMIT 1").Scan(&name)
		if typeErr != nil {
			return "There seems to have been a problem with the type you entered, please try again."
		}
	}
	//Name check
	if !skip && splitString[1] == "rarity" && len(splitString) > 2 {
		var nameErr = db.QueryRow("SELECT MULTIVERSEID FROM CARDS WHERE rarity like '" + splitString[2] + "%' ORDER BY RAND() LIMIT 1").Scan(&name)
		if nameErr != nil {
			return "There seems to have been a problem with the rarity you entered, please try again."
		}
	}

	//Default resonse if inproper command
	if name == "" {
		return "There seems to have been a problem with the command you entered, please try again."
	}

	db.Close()
	var url = fmt.Sprintf("http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=" + name + "&type=card")
	name = ""
	skip = false
	return url
}

//Char is used to generate potential character ideas
func Char() string {
	var adj string
	var race string
	var class string
	var location string
	var quirk string

	db, err = sql.Open("mysql", "root:1Merdenomz1@/magic")
	if err != nil {
		panic(err.Error())
	}

	//Queries need to be done like this due to bug in Mysql Driver for GO not supporting Stored Procedures
	_ = db.QueryRow("SELECT adjective FROM magic.character where adjective != '' ORDER BY RAND() LIMIT 1;").Scan(&adj)
	_ = db.QueryRow("SELECT race FROM magic.character where race != '' ORDER BY RAND() LIMIT 1;").Scan(&race)
	_ = db.QueryRow("SELECT class FROM magic.character where class != '' ORDER BY RAND() LIMIT 1;").Scan(&class)
	_ = db.QueryRow("SELECT location FROM magic.character where location != '' ORDER BY RAND() LIMIT 1;").Scan(&location)
	_ = db.QueryRow("SELECT quirk FROM magic.character where quirk != '' ORDER BY RAND() LIMIT 1;").Scan(&quirk)

	var chars = fmt.Sprintf("Your character is a " + adj + " " + race + ". They are a " + class + " from the " + location + ", that " + quirk)
	db.Close()
	return chars

}

func dbConnect() *sql.DB {
	db, err = sql.Open("mysql", "root:1Merdenomz1@/magic")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return db
}
