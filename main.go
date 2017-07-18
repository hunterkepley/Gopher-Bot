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
	// ALL COMMANDS: *hello", "*help", "*gopherify", "*gopher", "*roll"
	// Pictures of gophers:
	gophers = []string{"http://i.imgur.com/3tw6sII.jpg",
		"http://i.imgur.com/wUoSiNI.gif", "http://i.imgur.com/NfqwhN6.gif",
		"http://i.imgur.com/CBvD4d5.jpg", "http://i.imgur.com/CBMlinR.jpg",
		"http://i.imgur.com/32uPofb.jpg", "http://i.imgur.com/8jFGjsz.jpg",
		"http://i.imgur.com/seTJOPL.gif", "http://i.imgur.com/pBIh3pP.gif",
		"https://behrrake.files.wordpress.com/2008/04/fighing-gopher.jpg",
		"https://s-media-cache-ak0.pinimg.com/736x/63/af/78/63af782253e24944dd6d968acda29211--groundhog-pictures-happy-groundhog-day.jpg",
		"https://s-media-cache-ak0.pinimg.com/736x/1b/5a/2f/1b5a2fa52342a3cb980dbb38282683be--family-pictures-awkward-family-photos.jpg",
		"http://www.freakingnews.com/pictures/5500/Gopher--5651.jpg",
		"https://pics.me.me/i-could-spat-gopher-a-beer-funny-c3-15199885.png",
		"http://il8.picdn.net/shutterstock/videos/7339471/thumb/1.jpg",
		"http://images.gr-assets.com/books/1347514988l/14478480.jpg",
		"https://s-media-cache-ak0.pinimg.com/originals/19/80/1d/19801dba06ba8c5df1dff8cf64ef785c.jpg",
		"https://c1.staticflickr.com/3/2913/14753417043_7a92202e6a_b.jpg"}

	helpMsg = fmt.Sprintf("*hello -- Say hello\n*help -- Displays all commands\n*gopherify `YOUR MESSAGE` -- Gopherify's a message\n*gopher -- Display random gopher\n*roll `# #` -- Roll a random number\n*invite -- Displays invite link")

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
	if len(splitMsgLowered) > 0 {
		parseCommand(s, m, splitMsgLowered[0])
	}
}
