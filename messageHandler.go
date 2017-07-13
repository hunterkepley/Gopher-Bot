package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"math/rand"
	"strings"
	"strconv"
)

type MessageHandler struct { // This deals with all message inputs from the user.
	command string // Command user typed
	messageSlice []string // The message sent cut up
	commands []string // Commands: *Hello, *Help, *Gopherify
}

func (mH MessageHandler) BasicMessages(s *discordgo.Session, m *discordgo.MessageCreate) { // Deals with the basic commands
	if mH.command == mH.commands[0] { // Hello
		s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot`")
	}
	if mH.command == mH.commands[1] { // Help
		help := fmt.Sprintf("%s -- Say hello\n%s -- Displays all commands\n%s `YOUR MESSAGE` -- Gopherify's a message\n%s -- Display random gopher.\n%s `# #` -- Roll a random number", 
			mH.commands[0], mH.commands[1], mH.commands[2], mH.commands[3], mH.commands[4])
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "Commands:\n",
			Description: fmt.Sprintf("%s", help)})
	}
	if mH.command == mH.commands[2] { // Gopherify
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(mH.messageSlice)-1))) // -1 because the command is included.
	}
	if mH.command == mH.commands[4] { // Roll
		tNums := []string{"1", "100"}

		if(len(mH.messageSlice) >= 3) {
			tNumsTemp := []string{mH.messageSlice[1], mH.messageSlice[2]}
			tNums = tNumsTemp
		}

		num1, err := strconv.Atoi(tNums[0])
		if err != nil {
			num1 = 1
		}

		num2, err := strconv.Atoi(tNums[1])
		if err != nil {
			num2 = 100
		}

		rolledNumber := randInt(num1, num2)
		embedString := fmt.Sprintf("You rolled [%d - %d]:", num1, num2)

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: embedString, 
			Description: fmt.Sprintf("%d", rolledNumber)})
	}
}

func (mH MessageHandler) AdvancedMessages(s *discordgo.Session, m *discordgo.MessageCreate, gophers []string) {
	if(mH.command == mH.commands[3]) { // Gophers
		randGopher := randInt(0, len(gophers))
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{

			Title: "Gopher:", 

			Image: &discordgo.MessageEmbedImage{URL:gophers[randGopher]}})
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}