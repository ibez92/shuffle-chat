package discord

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffleChannelParticipants(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	ch, err := s.Channel(i.ChannelID)
	if err != nil {
		log.Fatal("shuffleChannelParticipants/Channel error: ", err)
	}

	userNames := []string{}
	for _, r := range ch.Recipients {
		if !r.Bot {
			userNames = append(userNames, r.Username)
		}
	}

	if len(userNames) > 0 {
		rand.Shuffle(len(userNames), func(i, j int) { userNames[i], userNames[j] = userNames[j], userNames[i] })
		return strings.Join(userNames, "\n")
	} else {
		return "No people in chat"
	}
}
