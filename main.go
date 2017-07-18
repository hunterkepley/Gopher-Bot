package main

import (
	"github.com/bwmarrin/discordgo"

	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Variables used for command line parameters
var (
	Token string
)

// Custom variables
var (
	helpMsg = fmt.Sprintf("Prefix: `*`\nhello -- Say hello\nhelp -- Displays all commands\nhelp `COMMAND` -- Displays extra help for a certain command\ngopherify `YOUR MESSAGE` -- Gopherify's a message\ngopher -- Display random gopher\nroll `# #` -- Roll a random number\ninvite -- Displays invite link\nbug `BUG INFO` -- Submit a bug to the creator of Gopher Bot")

	splitMsgLowered = []string{}
)

func makeSplitMessage(s *discordgo.Session, m *discordgo.MessageCreate) []string {
	// The message, split up
	splitMessage := strings.Fields(strings.ToLower(m.Content))

	return splitMessage
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Generate random seed for rng
	rand.Seed(time.Now().UTC().UnixNano())

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
	if m.Author.ID == s.State.User.ID {
		return
	}

	splitMsgLowered = makeSplitMessage(s, m)

	loadCommands()
	if len(splitMsgLowered) > 0 { // Prevented a really rare and weird bug about going out of index.
		parseCommand(s, m, splitMsgLowered[0]) // Really shouldnt happen since `MessageCreate` is about
	} // 										messages made on create...
}
