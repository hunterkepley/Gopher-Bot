package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func gopherCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	randGopher := randInt(0, len(gophers))
	var tGopherLink string
	var tGopherKey string
	if len(splitMsgLowered) > 1 {
		tGopherKey = splitMsgLowered[1]
		tGopherLink = searchGopher(tGopherKey)
	}
	if tGopherLink == "" {
		var ti int
		for _, v := range gophers {
			if ti == randGopher {
				tGopherLink = v
			} else {
				ti++
			}
		}
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{

		Title: fmt.Sprintf("Gopher %s:", tGopherKey),

		Image: &discordgo.MessageEmbedImage{URL: tGopherLink}})
}

func searchGopher(g string) (v string) { // ;)
	if gophers[g] == "" { // If it doesn't contain the key, then return a zero value
		return ""
	}
	return gophers[g] // Otherwise return the proper value
}
