package discord

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	discordFailedData = &discordgo.InteractionResponseData{
		Content: "Something went wrong",
	}
)

func (c *Client) shuffleChannelMembers(s discordSession, i *discordgo.InteractionCreate) {
	limit := 1000
	after := ""
	members := []*discordgo.Member{}
	for {
		respMembers, err := s.GuildMembers(i.GuildID, after, limit)
		if err != nil {
			fmt.Printf("shuffleChannelMembers/GuildMembers error: %v", err.Error())
			shuffleInteractionRespond(s, i, discordFailedData)
			break
		}

		if c.targetRole != "" {
			for _, m := range respMembers {
				for _, role := range m.Roles {
					if role == c.targetRole {
						members = append(members, m)
						break
					}
				}
			}
		} else {
			members = append(members, respMembers...)
		}

		after = respMembers[len(respMembers)-1].User.ID
		if len(respMembers) < limit {
			break
		}
	}

	userNames := []string{}
	for _, m := range members {
		if !m.User.Bot {
			userNames = append(userNames, m.User.Username)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(userNames), func(i, j int) { userNames[i], userNames[j] = userNames[j], userNames[i] })

	data := &discordgo.InteractionResponseData{
		Content: strings.Join(userNames, "\n"),
	}
	shuffleInteractionRespond(s, i, data)
}

func shuffleInteractionRespond(s discordSession, i *discordgo.InteractionCreate, data *discordgo.InteractionResponseData) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})

	if err != nil {
		fmt.Printf("shuffleInteractionRespond/InteractionRespond error: %v", err.Error())
		//nolint:errcheck
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: discordFailedData,
		})
	}
}
