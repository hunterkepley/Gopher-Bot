package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"strings"
)

type MessageHandler struct {
	command string // Command user typed
	messageSlice []string
	commands []string // Commands: *Hello, *Help, *Gopherify
}

func NewMessageHandler(cmd string, messages []string, cmds []string) *MessageHandler { // 'Constructor' for MessageHandler
	return &MessageHandler{cmd, messages, cmds}
}

func (mH MessageHandler) BasicMessages(s *discordgo.Session, m *discordgo.MessageCreate) { // Deals with the basic commands
	if mH.command == mH.commands[0] { // Hello
		s.ChannelMessageSend(m.ChannelID, "`Henlo, I am Gopher Bot`")
	}
	if mH.command == mH.commands[1] { // Help
		help := fmt.Sprintf("`Commands:\n%s\n%s\n%s <YOUR MESSAGE>`", mH.commands[0], mH.commands[1], mH.commands[2])
		s.ChannelMessageSend(m.ChannelID, help)
	}
	if mH.command == mH.commands[2] { // Gopherify
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(mH.messageSlice)-1))) // -1 because the command is included.
	}
}