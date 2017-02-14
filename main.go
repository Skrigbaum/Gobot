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
	if msg == "vyrtec" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "The Autist?")
	}

	if msg == "sly" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Weaboo for sure")
	}

	// If the message is "ping" reply with "Pong!"
	if msg == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if msg == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if strings.Contains(msg, "!card") {
		cardName := Mtg()
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
func Mtg() string {
	var name string

	db, err = sql.Open("mysql", "root:1Merdenomz1@/magic")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.QueryRow("SELECT MULTIVERSEID FROM CARDS ORDER BY RAND() LIMIT 1").Scan(&name)
	if err != nil {
		panic(err.Error())
	}

	//err = card.Scan(&cardName)
	if err != nil {
		panic(err.Error())
	}
	db.Close()
	var url = fmt.Sprintf("http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=" + name + "&type=card")
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
