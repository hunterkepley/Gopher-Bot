package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"strings"
)

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot, I have random utilities and fun features. My prefix is `*`. Github: https://www.github.com/hunterkepley/Gopher-Bot`. Message `TheVariant#9315` if the bot is down or has a bug!")
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
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Send a bug via:",
		Description: "My github:\thttps://www.github.com/hunterkepley/Gopher-Bot\nMy email:\t`kepley.l.hunter@gmail.com`\nMy Discord Name:\t`TheVariant#9315`"})
}
