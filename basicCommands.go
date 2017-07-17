package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"strings"
)

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot, I have random utilities and fun features. Github: https://www.github.com/hunterkepley/gopher-bot`")
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Commands:\n",
		Description: fmt.Sprintf("%s", helpMsg)})
}

func gopherifyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(splitMsgLowered)-1))) // -1 because the command is included.
}

func inviteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Invite me!:\n",
		Description: "Click this link:\nhttps://discordapp.com/oauth2/authorize?client_id=334056784748609547&scope=bot"})
}
