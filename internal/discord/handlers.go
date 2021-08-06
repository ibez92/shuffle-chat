package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) readyHandler(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateGameStatus(1, "Shuffler")
	if err != nil {
		log.Fatal("UpdateStatusComplex error: ", err)
	}
}

func (c *Client) commandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		shuffleCommand: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			c.shuffleChannelMembers(s, i)
		},
	}

	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
