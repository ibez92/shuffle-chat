package discord

import "github.com/bwmarrin/discordgo"

type (
	discordSession interface {
		GuildMembers(string, string, int) ([]*discordgo.Member, error)
		InteractionRespond(*discordgo.Interaction, *discordgo.InteractionResponse) error
	}
)
