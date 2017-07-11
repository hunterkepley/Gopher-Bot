package main

import (
	"github.com/bwmarrin/discordgo"

	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Variables used for command line parameters
var (
	Token string
)

// Custom variables
var (
	commands []string = []string{"*Hello", "*Help", "*Gopherify"} // Commands
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord sessions using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to stop")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session
	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) { // Message handling
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but if it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// The message, split up
	splitMessage := strings.Split(m.Content, " ")

	msgHandler := NewMessageHandler(splitMessage[0], splitMessage, commands) // splitMessage[0] is the command itself

	// Handles basic messages
	msgHandler.BasicMessages(s, m)

}

