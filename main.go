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
	buffer  = make([][]byte, 0)
	Token   string
	BotID   string
	err     error
	message bytes.Buffer
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

	//if the message is !gobot return a helpful message
	if msg == "!gobot" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `Hi I'm Gobot. Currently typing '!card returns a random magic card',
'!card type *X* returns a card of that type', '!card rarity *X* returns a card of that rarity'.
'!fantasy -c returns a random D&D character', and '!fantasy - preturns a random place name.'
'!name -m returns a random male name', '!name -f returns a random female name', '!name -l to return a family name'
Please excuse any of my syntax or grammar errors, my creator doesn't know how to spell well.`)
	}

	//card flag section
	if strings.Contains(msg, "!card") {
		cardName := functions.Card(msg)
		_, _ = s.ChannelMessageSend(m.ChannelID, cardName)
	}

	//fantasy flags
	//fantasy Fantasy flag
	if strings.Contains(msg, "!fantasy") {
		placeName := functions.Fantasy(msg)

		_, _ = s.ChannelMessageSend(m.ChannelID, placeName)
	}
	//Random fantasy character flag
	if strings.Contains(msg, "!char") {
		character := functions.Char()
		if err != nil {
			fmt.Println("Error retriving character,", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Stop breaking stuff.")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, character)

	}

	//Character Name flag
	if strings.Contains(msg, "!name") {
		var splitString = strings.Split(msg, " ")
		placeName := functions.Name(splitString)

		_, _ = s.ChannelMessageSend(m.ChannelID, placeName)
	}

	//Leauge of legends flags
	//League summoner recent game history flag
	if strings.Contains(msg, "!recent") {
		var game = functions.League(msg)
		if err != nil {
			fmt.Println("Error retriving league info,", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Stop breaking stuff.")
			return
		}
		for _, resp := range game {
			message.WriteString(resp + "\n")
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, message.String())
		message.Reset()
	}
	//Random League Character generation
	if strings.Contains(msg, "!gorandom") {
		randomChar := functions.Random()
		if err != nil {
			fmt.Println("Error retriving random character,", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Stop breaking stuff.")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, randomChar)

	}
}
