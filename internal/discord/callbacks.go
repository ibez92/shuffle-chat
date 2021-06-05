package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Open(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return dg, nil
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateGameStatus(0, "!let's shuffle some conf")
	if err != nil {
		log.Fatal(err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "!shuffle-zoom-conf" {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if c.Name != "bot" {
		return
	}
}
