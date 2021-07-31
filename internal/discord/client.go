package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	Session *discordgo.Session
}

func NewClient(token, guildID string) *Client {
	c := Client{}
	session, err := c.openDiscordSession(token, guildID)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	c.Session = session

	return &c
}

func (c *Client) openDiscordSession(token, guildID string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.AddHandler(readyHandler)
	session.AddHandler(commandsHandler)

	// Open the websocket and begin listening.
	err = session.Open()
	if err != nil {
		return nil, err
	}

	createApplicationCommands(session, guildID)

	return session, nil
}

func createApplicationCommands(session *discordgo.Session, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        shuffleCommand,
			Description: "Shuffle participants of current text channel",
		},
	}

	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
