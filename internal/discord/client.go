package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	Session      *discordgo.Session
	targetRoleID string
}

func NewClient(token, targetRoleID string) *Client {
	c := Client{targetRoleID: targetRoleID}
	session, err := c.openDiscordSession(token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	c.Session = session

	return &c
}

func (c *Client) openDiscordSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.AddHandler(c.readyHandler)
	session.AddHandler(c.commandsHandler)

	// Open the websocket and begin listening.
	err = session.Open()
	if err != nil {
		return nil, err
	}

	createApplicationCommands(session)

	return session, nil
}

func createApplicationCommands(session *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        shuffleCommand,
			Description: "Shuffle participants of current text channel",
			Options: []*discordgo.ApplicationCommandOption{
				// {
				// 	Name:        "role-option",
				// 	Type:        discordgo.ApplicationCommandOptionString,
				// 	Description: "Filter members by role",
				// 	Required:    false,
				// },
				{
					Name:        "role-option",
					Type:        discordgo.ApplicationCommandOptionRole,
					Description: "Filter members by role",
					Required:    false,
				},
			},
		},
	}

	for _, v := range commands {
		for _, g := range session.State.Guilds {
			_, err := session.ApplicationCommandCreate(session.State.User.ID, g.ID, v)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", v.Name, err)
			}
		}
	}
}
