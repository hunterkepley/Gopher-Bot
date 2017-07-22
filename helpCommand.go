package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
