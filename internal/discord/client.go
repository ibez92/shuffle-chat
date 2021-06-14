package discord

import (
	"log"
	"shufflezoommeeting/internal/zoom"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	Ds         *discordgo.Session
	zoomClient *zoom.Client
}

func NewClient(token string, zoomClient *zoom.Client) *Client {
	c := Client{zoomClient: zoomClient}
	ds, err := c.openDiscordSession(token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	c.Ds = ds

	return &c
}

func (c *Client) openDiscordSession(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate(c.zoomClient))

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return dg, nil
}
