package main

import "github.com/bwmarrin/discordgo"

func inviteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Invite me!:",
		Description: "Click this link:\nhttps://discordapp.com/oauth2/authorize?client_id=334056784748609547&scope=bot"})
}
