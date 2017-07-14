package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
)

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot`")
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Commands:\n",
		Description: fmt.Sprintf("%s", helpMsg)})
}