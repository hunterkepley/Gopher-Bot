package main

import (
	"github.com/bwmarrin/discordgo"
)

func gopherCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	randGopher := randInt(0, len(gophers))
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{

		Title: "Gopher:",

		Image: &discordgo.MessageEmbedImage{URL: gophers[randGopher]}})
}
