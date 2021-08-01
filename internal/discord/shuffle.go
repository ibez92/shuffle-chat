package discord

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const noPeopleResult = "No people in chat"

func (c *Client) shuffleChannelParticipants(recipients []*discordgo.User, members []*discordgo.Member) string {
	usersWithRole := []*discordgo.User{}
	for _, m := range members {
		for _, role := range m.Roles {
			if role == c.targetRole {
				usersWithRole = append(usersWithRole, m.User)
				break
			}
		}
	}
	if len(usersWithRole) == 0 {
		return noPeopleResult
	}

	users := []*discordgo.User{}
	for _, r := range recipients {
		for _, u := range usersWithRole {
			if u.ID == r.ID && !r.Bot {
				users = append(users, r)
				break
			}
		}
	}

	if len(users) == 0 {
		return noPeopleResult
	}

	userNames := []string{}
	for _, u := range users {
		userNames = append(userNames, u.Username)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(userNames), func(i, j int) { userNames[i], userNames[j] = userNames[j], userNames[i] })
	return strings.Join(userNames, "\n")
}
