package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
