package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"fmt"
	"math/rand"
	"strconv"
)

func gopherCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	randGopher := randInt(0, len(gophers))
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{

		Title: "Gopher:",

		Image: &discordgo.MessageEmbedImage{URL: gophers[randGopher]}})
}

func randInt(min int, max int) int {
	if max == min {
		return max
	}
	if min > max {
		return max + rand.Intn(min-max)
	}
	return min + rand.Intn(max-min)
}

func rollCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	tNums := []string{"1", "100"}

	if len(splitMsgLowered) >= 3 {
		tNumsTemp := []string{splitMsgLowered[1], splitMsgLowered[2]}
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
		Title:       embedString,
		Description: fmt.Sprintf(":game_die: : %d", rolledNumber)})
}

func bugCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(splitMsgLowered) > 1 {
		bugReport := strings.Join(splitMsgLowered[1:], " ")
		channel, err := s.UserChannelCreate("121105861539135490")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to send bug report, user not found. Please go on the github and make an `issue` about this. Use `*hello` to see the github.")
		}
		authorChannel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to message user who sent bug report. Please go on the github and make an `issue` about this. Use `*hello` to see the github.")
		}
		s.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
			Title: "Bug Report:",
			Description: fmt.Sprintf("<@!%s>: %s",
				m.Author.ID,
				bugReport)})
		s.ChannelMessageSend(authorChannel.ID, "Sent bug report! Thanks for helping :smile:")
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "You need to write a message after `*bug`!"})
	}
}

func gopherRequestCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(splitMsgLowered) == 2 {
		gopherRequest := splitMsgLowered[1]
		channel, err := s.UserChannelCreate("121105861539135490")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to send gopher request, user not found. Please go on the github and make an `issue` about this. Use `*hello` to see the github.")
		}
		authorChannel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to message user who sent gopher request. Please go on the github and make an `issue` about this. Use `*hello` to see the github.")
		}
		s.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
			Title: "Gopher Request:",
			Description: fmt.Sprintf("<@!%s>: %s",
				m.Author.ID,
				gopherRequest)})
		s.ChannelMessageSend(authorChannel.ID, "Sent gopher request!")
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "You need to paste a link after `*gopher`!"})
	}
}
