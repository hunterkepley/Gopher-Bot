package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"
)

var (
	commMap = make(map[string]Command)

	hello     = Command{"*hello", "Says hello", helloCommand}
	help      = Command{"*help", "Displays all commands", helpCommand}
	gopherify = Command{"*gopherify", "Gopherify's a message", gopherifyCommand}
	gopher    = Command{"*gopher", "Displays random gopher", gopherCommand}
	roll      = Command{"*roll", "Rolls a random number from x to x", rollCommand}
	invite    = Command{"*invite", "Displays invite link", inviteCommand}
)

// Command : Every command is made into a struct to make it simpler to work with and eliminate if statements
type Command struct {
	name        string
	description string
	exec        func(*discordgo.Session, *discordgo.MessageCreate)
}

func loadCommands() {
	commMap[hello.name] = hello
	commMap[help.name] = help
	commMap[gopherify.name] = gopherify
	commMap[gopher.name] = gopher
	commMap[roll.name] = roll
	commMap[invite.name] = invite
}

func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	if command == strings.ToLower(commMap[command].name) {
		commMap[command].exec(s, m)
	}
	return
}
