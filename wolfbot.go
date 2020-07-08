package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Wolfbot contains the Discord token and session object.
// All bot functionality is implemented within it.
type Wolfbot struct {
	token   string
	session *discordgo.Session
}

// Configure reads in the config.yaml file and configures the WolfBot struct.
// Configure MUST be called before Connect or Run
func (wb *Wolfbot) Configure() {
	// var err error
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("No config file found at config/config.yaml.", err)
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Fatal error reading config file: %s \n", err)
		}
	}

	// TODO: Connect to DB here.

	wb.token = viper.GetString("discordToken")
}

// Connect establishes the websocket connection to Discord, authenticates the bot and populates the session property.
// Configure MUST be called before Connect.
// Connect MUST be called before Run.
func (wb *Wolfbot) Connect() {
	var err error

	if wb.token == "" {
		log.Fatal("Missing discord token in config/config.yaml.")
	}

	wb.session, err = discordgo.New("Bot " + wb.token)
	if err != nil {
		log.Fatalf("Fatal error creating discord bot: %s.\n", err)
	}

	wb.session.AddHandler(wb.messageCreateHandler)
}

// Run begins listening on the websocket connection.
// Configure and Connect MUST be called before Run.
func (wb *Wolfbot) Run() {
	// Open the websocket and begin listening.
	err := wb.session.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}
}

// Stop cleanly closes down the Discord session.
func (wb *Wolfbot) Stop() {
	wb.session.Close()
}

// messageCreateHandler handles an incoming Discord message.
func (wb *Wolfbot) messageCreateHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// TODO: This is only for demonstration, remove this so we're not logging every discord message...
	log.Printf("%s: %s\n", message.Author.Username, message.ContentWithMentionsReplaced())

}
