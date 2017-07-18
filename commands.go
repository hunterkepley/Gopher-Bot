package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"
)

var (
	commMap = make(map[string]Command)

	hello     = Command{"hello", "Says hello, just an introduction, nothing more, nothing less.", helloCommand}
	help      = Command{"help", "Displays all commands, pretty obvious. Also can display specific information using `*help` and a command after, for example, `*help gopher`.", helpCommand}
	gopherify = Command{"gopherify", "Gopherify's a message, basically just squeaks at ya.", gopherifyCommand}
	gopher    = Command{"gopher", "Displays random gopher out of a pretty large selection randomly.", gopherCommand}
	roll      = Command{"roll", "Rolls a random number from x to x, automatically defaults to 1 - 100 if you mess up or make too large of a number.", rollCommand}
	invite    = Command{"invite", "Displays invite link to invite to other servers.", inviteCommand}
	bug       = Command{"bug", "Send a bug by typing `*bug This is broken, here's how to reprodruce it` to the creator of Gopher-Bot!", bugCommand}
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
	commMap[bug.name] = bug
}

func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	if strings.Contains(string(command[0]), "*") { // If the first word of the message starts with `*`:`
		command = string(command[1:]) // Remove the `*` from the command
		if command == strings.ToLower(commMap[command].name) {
			commMap[command].exec(s, m)
		}
	}
	return
}
