package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func readyHandler(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateGameStatus(-1, "Let's shuffle this")
	if err != nil {
		log.Fatal("UpdateStatusComplex error: ", err)
	}
}

func commandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		shuffleCommand: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: shuffleChannelParticipants(s, i),
				},
			})
		},
	}

	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}