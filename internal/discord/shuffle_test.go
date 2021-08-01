package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func Test_shuffle(t *testing.T) {
	targetRole := "tr"

	tests := []struct {
		name       string
		recipients []*discordgo.User
		members    []*discordgo.Member
		result     string
	}{
		{
			name: "Success",
			recipients: []*discordgo.User{
				{ID: "ID1", Username: "1"},
				{ID: "ID2", Username: "2"},
				{ID: "ID3", Username: "3"},
				{ID: "ID4", Username: "4"},
			},
			members: []*discordgo.Member{
				{
					User:  &discordgo.User{ID: "ID1"},
					Roles: []string{targetRole},
				},
				{
					User:  &discordgo.User{ID: "ID2"},
					Roles: []string{"invalidRole"},
				},
				{
					User:  &discordgo.User{ID: "invalidID"},
					Roles: []string{targetRole},
				},
				{
					User:  &discordgo.User{ID: "ID4"},
					Roles: []string{targetRole},
				},
			},
			result: "1\n4",
		},
		{
			name:       "No recipients",
			recipients: []*discordgo.User{},
			members: []*discordgo.Member{
				{
					User:  &discordgo.User{ID: "ID1"},
					Roles: []string{targetRole},
				},
				{
					User:  &discordgo.User{ID: "ID2"},
					Roles: []string{"invalidRole"},
				},
			},
			result: noPeopleResult,
		},
		{
			name: "No Members",
			recipients: []*discordgo.User{
				{ID: "ID1", Username: "1"},
			},
			members: []*discordgo.Member{},
			result:  noPeopleResult,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := Client{targetRole: targetRole}
			got := c.shuffleChannelParticipants(tt.recipients, tt.members)
			assert.Equal(t, tt.result, got)
		})
	}
}
