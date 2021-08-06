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
	trID := c.targetRoleID
	opts := i.ApplicationCommandData().Options
	if len(opts) > 0 {
		trID = opts[0].RoleValue(c.Session, i.GuildID).ID
	}

	ch, err := s.Channel(i.ChannelID)
	if err != nil {
		fmt.Printf("shuffleChannelMembers/Channel error: %v\n", err.Error())
		shuffleInteractionRespond(s, i, discordFailedData)
		return
	}

	chWithRole := false
	for _, p := range ch.PermissionOverwrites {
		if p.Type == 0 && p.ID == trID {
			chWithRole = true
		}

		if chWithRole {
			break
		}
	}

	if !chWithRole {
		data := &discordgo.InteractionResponseData{
			Content: "No role in this channel",
		}
		shuffleInteractionRespond(s, i, data)
		return
	}

	userNames := []string{}
	limit := 1000
	after := ""
	for {
		respMembers, err := s.GuildMembers(i.GuildID, after, limit)
		if err != nil {
			fmt.Printf("shuffleChannelMembers/GuildMembers error: %v\n", err.Error())
			shuffleInteractionRespond(s, i, discordFailedData)
			break
		}

		for _, m := range respMembers {
			if m.User.Bot {
				continue
			}

			for _, role := range m.Roles {
				if role == trID {
					userNames = append(userNames, m.User.Mention())
					break
				}
			}
		}

		after = respMembers[len(respMembers)-1].User.ID
		if len(respMembers) < limit {
			break
		}
	}

	if len(userNames) == 0 {
		data := &discordgo.InteractionResponseData{
			Content: "No members found",
		}
		shuffleInteractionRespond(s, i, data)
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(userNames), func(i, j int) { userNames[i], userNames[j] = userNames[j], userNames[i] })

	data := &discordgo.InteractionResponseData{
		Content: strings.Join(userNames, "\n"),
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
		},
	}
	shuffleInteractionRespond(s, i, data)
}

func shuffleInteractionRespond(s discordSession, i *discordgo.InteractionCreate, data *discordgo.InteractionResponseData) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})

	if err != nil {
		fmt.Printf("shuffleInteractionRespond/InteractionRespond error: %v\n", err.Error())
		//nolint:errcheck
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: discordFailedData,
		})
	}
}
