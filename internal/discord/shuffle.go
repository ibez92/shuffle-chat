package discord

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const noPeopleResult = "No people in chat"

func shuffleChannelParticipants(recipients []*discordgo.User) string {
	userNames := []string{}
	for _, r := range recipients {
		if !r.Bot {
			userNames = append(userNames, r.Username)
		}
	}

	if len(userNames) > 0 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(userNames), func(i, j int) { userNames[i], userNames[j] = userNames[j], userNames[i] })
		return strings.Join(userNames, "\n")
	}

	return noPeopleResult
}
