package discord

import "github.com/bwmarrin/discordgo"

type (
	discordSession interface {
		Guild(string) (*discordgo.Guild, error)
		Channel(string) (*discordgo.Channel, error)
		GuildMembers(string, string, int) ([]*discordgo.Member, error)
		InteractionRespond(*discordgo.Interaction, *discordgo.InteractionResponse) error
	}
)
