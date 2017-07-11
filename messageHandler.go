package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"math/rand"
	"strings"
)

type MessageHandler struct { // This deals with all message inputs from the user.
	command string // Command user typed
	messageSlice []string
	commands []string // Commands: *Hello, *Help, *Gopherify
}

func (mH MessageHandler) BasicMessages(s *discordgo.Session, m *discordgo.MessageCreate) { // Deals with the basic commands
	if mH.command == mH.commands[0] { // Hello
		s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot`")
	}
	if mH.command == mH.commands[1] { // Help
		help := fmt.Sprintf("`Commands:\n%s\n%s\n%s <YOUR MESSAGE>\n%s`", mH.commands[0], mH.commands[1], mH.commands[2], mH.commands[3])
		s.ChannelMessageSend(m.ChannelID, help)
	}
	if mH.command == mH.commands[2] { // Gopherify
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(mH.messageSlice)-1))) // -1 because the command is included.
	}
}

func (mH MessageHandler) AdvancedMessages(s *discordgo.Session, m *discordgo.MessageCreate, gophers []string) {
	if(mH.command == mH.commands[3]) { // Gopher
		randGopher := randInt(0, len(gophers))
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "Gopher:", Image: &discordgo.MessageEmbedImage{URL:gophers[randGopher]}})
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}