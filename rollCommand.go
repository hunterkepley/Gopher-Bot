package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func randInt(min int, max int) int {
	if max == min {
		return max
	}
	if min > max {
		return max + rand.Intn(min-max)
	}
	return min + rand.Intn(max-min)
}

func rollCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	tNums := []string{"1", "100"}

	if len(splitMsgLowered) >= 3 {
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
		Title:       embedString,
		Description: fmt.Sprintf(":game_die: : %d", rolledNumber)})
}
