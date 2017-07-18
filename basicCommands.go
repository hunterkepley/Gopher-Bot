package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"strings"
)

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "`Henlo, I am Gopher Bot, I have random utilities and fun features.`",
		Description: "My prefix is `*`. Github: https://www.github.com/hunterkepley/Gopher-Bot. Message `TheVariant#9315` if the bot is down and use `*bug` to see how to submit a bug!"})
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(splitMsgLowered) == 1 {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Commands:",
			Description: fmt.Sprintf("%s", helpMsg)})
	} else {
		if splitMsgLowered[1] == strings.ToLower(commMap[splitMsgLowered[1]].name) {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s Help:", splitMsgLowered[1]),
				Description: commMap[splitMsgLowered[1]].description})
		} else {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "Command does not exist!"})
		}
	}

}

func gopherifyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(splitMsgLowered)-1))) // -1 because the command is included.
}

func inviteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Invite me!:",
		Description: "Click this link:\nhttps://discordapp.com/oauth2/authorize?client_id=334056784748609547&scope=bot"})
}

func bugCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(splitMsgLowered) > 1 {
		bugReport := strings.Join(splitMsgLowered[1:], " ")
		channel, err := s.UserChannelCreate("121105861539135490")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to send bug report, user not found. Please go on the github and make an `issue` about this. Use `*hello` to see the github.")
		}

		s.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
			Title: "Bug Report:",
			Description: fmt.Sprintf("<@!%s>: %s",
				m.Author.ID,
				bugReport)})
		s.ChannelMessageSend(m.Author.ID, "Sent bug report! Thanks for helping :smile:")
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "You need to write a message after `*bug`!"})
	}
}
