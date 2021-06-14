package discord

import (
	"fmt"
	"log"
	"math/rand"
	"shufflezoommeeting/internal/zoom"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateGameStatus(0, "!let's shuffle some meetings")
	if err != nil {
		log.Fatal("UpdateGameStatus error: ", err)
	}
}

func messageCreate(zc *zoom.Client) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.LastIndex(m.Content, "!shuffle-zoom-meeting") != 0 {
			return
		}

		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			fmt.Println("Channel error: ", err)
			return
		}

		if c.Name != "bot" {
			return
		}

		mID := ""
		parts := strings.Split(m.Content, " ")
		if len(parts) > 1 {
			mID = parts[1]
		}

		ps, err := zc.GetMeetingParticipants(mID)
		if err != nil {
			fmt.Println("GetMeetingParticipants error: ", err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })

		names := make([]string, len(ps))
		for i, p := range ps {
			names[i] = p.UserName
		}
		s.ChannelMessageSend(c.ID, strings.Join(names, "\n"))
	}
}
