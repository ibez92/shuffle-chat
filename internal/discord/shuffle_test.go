package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func Test_shuffle(t *testing.T) {
	tests := []struct {
		name       string
		recipients []*discordgo.User
		result     string
	}{
		{
			name: "Success",
			recipients: []*discordgo.User{
				{Username: "1"},
				{Username: "2"},
			},
			result: "1\n2",
		},
		{
			name:       "No recipients",
			recipients: []*discordgo.User{},
			result:     noPeopleResult,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffleChannelParticipants(tt.recipients)
			assert.Equal(t, tt.result, got)
		})
	}
}
