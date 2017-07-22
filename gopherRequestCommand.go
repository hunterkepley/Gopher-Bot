package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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
