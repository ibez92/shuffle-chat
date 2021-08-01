package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) readyHandler(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateGameStatus(-1, "Let's shuffle this")
	if err != nil {
		log.Fatal("UpdateStatusComplex error: ", err)
	}
}

func (c *Client) commandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		shuffleCommand: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			guild, err := s.Guild(i.GuildID)
			if err != nil {
				log.Fatal("commandsHandler/guild error: ", err)
			}

			ch, err := s.Channel(i.ChannelID)
			if err != nil {
				log.Fatal("commandsHandler/channel error: ", err)
			}

			content := c.shuffleChannelParticipants(ch.Recipients, guild.Members)
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
			if err != nil {
				fmt.Printf("discord/handlers/"+shuffleCommand+" failed. Error: %v", err)
			}
		},
	}

	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
