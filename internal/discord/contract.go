package discord

import "github.com/bwmarrin/discordgo"

type (
	DiscordSessionInf interface {
		Channel(string) (*discordgo.Channel, error)
	}
)
