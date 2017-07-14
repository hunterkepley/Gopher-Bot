package main

import (
	"github.com/bwmarrin/discordgo"
	
	"fmt"
	"strings"
	"strconv"
	"math/rand"
)

func gopherCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	randGopher := randInt(0, len(gophers))
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{

		Title: "Gopher:", 

		Image: &discordgo.MessageEmbedImage{URL:gophers[randGopher]}})
}

func randInt(min int, max int) int {
	if(max == min) {
		return max
	}
	if(min > max) {
		return max + rand.Intn(min-max)
	}
	return min + rand.Intn(max-min)
}

func gopherifyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", strings.Repeat("Squeak ", len(splitMsgLowered)-1))) // -1 because the command is included.
}

func rollCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
		tNums := []string{"1", "100"}

		if(len(splitMsgLowered) >= 3) {
			tNumsTemp := []string{splitMsgLowered[1], splitMsgLowered[2]}
			tNums = tNumsTemp
		}

		num1, err := strconv.Atoi(tNums[0])
		if err != nil {
			num1 = 1
		}

		num2, err := strconv.Atoi(tNums[1])
		if err != nil {
			num2 = 100
		}

		rolledNumber := randInt(num1, num2)
		embedString := fmt.Sprintf("You rolled [%d - %d]:", num1, num2)

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: embedString, 
			Description: fmt.Sprintf(":game_die: : %d", rolledNumber)})
}