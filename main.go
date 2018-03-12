package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"github.com/Skrigbaum/Gobot/functions"
	"github.com/Skrigbaum/Gobot/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	buffer    = make([][]byte, 0)
	Token     string
	BotID     string
	err       error
	message   bytes.Buffer
	switchCon []string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	models.LoadConfig()

}

func main() {

	//Init DB connection
	models.InitDB()

	// Create a new Discord session using the provided bot token.
	dg, errCon := discordgo.New("Bot " + Token)
	if errCon != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, usrErr := dg.User("@me")
	if usrErr != nil {
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
	switchCon = strings.Fields(msg)

	switch switchCon[0] {
	case "!card":
		cardName := functions.Card(msg)
		_, _ = s.ChannelMessageSend(m.ChannelID, cardName)
	case "!gobot":
		_, _ = s.ChannelMessageSend(m.ChannelID, `Hi I'm Gobot. Currently typing '!card returns a random magic card',
			'!card type *X* returns a card of that type', '!card rarity *X* returns a card of that rarity'.
			'!fantasy -c returns a random D&D character', and '!fantasy - p returns a random place name, and !fantasy -a generates an adventure idea.'
			'!name -m returns a random male name', '!name -f returns a random female name', '!name -l to return a family name',
			'/roll XdY will also roll die for you! ex. /roll 3d6. Feel free to add modifiers to the dice as well. ex. 3d6+2'
			Please excuse any of my syntax or grammar errors, my creator doesn't know how to spell well.`)
	case "!load":
		fmt.Println("Hit the Load")
		functions.LoadCards()
		_, _ = s.ChannelMessageSend(m.ChannelID, "Success")
	case "!set":
		setName := functions.SetName(msg)
		_, _ = s.ChannelMessageSend(m.ChannelID, setName)
	case "/roll":
		roll := functions.Rolling(msg)
		_, _ = s.ChannelMessageSend(m.ChannelID, roll)
	case "!fantasy":
		placeName := functions.Fantasy(msg)
		_, _ = s.ChannelMessageSend(m.ChannelID, placeName)
	case "!char":
		character := functions.Char()
		if err != nil {
			fmt.Println("Error retriving character,", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Stop breaking stuff.")
			return
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, character)
	case "!name":
		var splitString = strings.Split(msg, " ")
		placeName := functions.Name(splitString)
		_, _ = s.ChannelMessageSend(m.ChannelID, placeName)
	}
}
