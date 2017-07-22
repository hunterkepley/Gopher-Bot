package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func gopherifyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(splitMsgLowered)-1))) // -1 because the command is included
}
