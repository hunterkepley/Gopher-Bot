package main

import (
	"github.com/bwmarrin/discordgo"
)

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "`Henlo, I am Gopher Bot, I have random utilities and fun features.`",
		Description: "My prefix is `*`. Github: https://www.github.com/hunterkepley/Gopher-Bot. Message `TheVariant#9315` if the bot is down and use `*bug` to see how to submit a bug!"})
}
