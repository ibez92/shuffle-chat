package discord

// TODO: Try to test this
// import (
// 	"testing"

// 	"github.com/bwmarrin/discordgo"
// 	"github.com/golang/mock/gomock"

// 	m "github.com/psy1992/shuffle-chat/internal/discord/mocks"
// )

// //nolint:funlen
// func Test_shuffleChannelMembers(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	guildID := "guildID"
// 	chID := "chID"

// 	type discordContract struct {
// 		guildID string
// 		members []*discordgo.Member
// 		ic      *discordgo.InteractionCreate
// 		ch      *discordgo.Channel
// 		mErr    error
// 		iErr    error
// 		chErr   error
// 	}

// 	tests := []struct {
// 		name            string
// 		discordContract discordContract
// 		tr              string
// 		result          string
// 	}{
// 		{
// 			name: "Success",
// 			discordContract: discordContract{
// 				guildID: guildID,
// 				members: []*discordgo.Member{
// 					{
// 						User: &discordgo.User{ID: "ID1", Username: "U1"},
// 					},
// 				},
// 				ic: &discordgo.InteractionCreate{
// 					Interaction: &discordgo.Interaction{
// 						GuildID:   guildID,
// 						ChannelID: chID,
// 					},
// 				},
// 				ch: &discordgo.Channel{
// 					PermissionOverwrites: []*discordgo.PermissionOverwrite{
// 						{
// 							ID:   "ID1",
// 							Type: 1,
// 						},
// 					},
// 				},
// 			},
// 			result: "U1",
// 		},
// 		{
// 			name: "With role filter",
// 			discordContract: discordContract{
// 				guildID: guildID,
// 				members: []*discordgo.Member{
// 					{
// 						User:  &discordgo.User{ID: "ID1", Username: "U1"},
// 						Roles: []string{"Role1", "Role2"},
// 					},
// 					{
// 						User:  &discordgo.User{ID: "ID2", Username: "U2"},
// 						Roles: []string{"Role2", "Role3"},
// 					},
// 				},
// 				ic: &discordgo.InteractionCreate{
// 					Interaction: &discordgo.Interaction{
// 						GuildID:   guildID,
// 						ChannelID: chID,
// 					},
// 				},
// 				ch: &discordgo.Channel{
// 					PermissionOverwrites: []*discordgo.PermissionOverwrite{
// 						{
// 							ID:   "ID1",
// 							Type: 1,
// 						},
// 					},
// 				},
// 			},
// 			tr:     "Role1",
// 			result: "U1",
// 		},
// 		{
// 			name: "It must skip bots",
// 			discordContract: discordContract{
// 				guildID: guildID,
// 				members: []*discordgo.Member{
// 					{
// 						User: &discordgo.User{ID: "ID1", Username: "U1", Bot: true},
// 					},
// 					{
// 						User: &discordgo.User{ID: "ID2", Username: "U2"},
// 					},
// 				},
// 				ic: &discordgo.InteractionCreate{
// 					Interaction: &discordgo.Interaction{
// 						GuildID:   guildID,
// 						ChannelID: chID,
// 					},
// 				},
// 				ch: &discordgo.Channel{
// 					PermissionOverwrites: []*discordgo.PermissionOverwrite{
// 						{
// 							ID:   "ID1",
// 							Type: 1,
// 						},
// 						{
// 							ID:   "ID2",
// 							Type: 1,
// 						},
// 					},
// 				},
// 			},
// 			result: "U2",
// 		},
// 		{
// 			name: "It filters by permissions",
// 			discordContract: discordContract{
// 				guildID: guildID,
// 				members: []*discordgo.Member{
// 					{
// 						User: &discordgo.User{ID: "ID1", Username: "U1"},
// 					},
// 					{
// 						User: &discordgo.User{ID: "ID2", Username: "U2"},
// 					},
// 				},
// 				ic: &discordgo.InteractionCreate{
// 					Interaction: &discordgo.Interaction{
// 						GuildID:   guildID,
// 						ChannelID: chID,
// 					},
// 				},
// 				ch: &discordgo.Channel{
// 					PermissionOverwrites: []*discordgo.PermissionOverwrite{
// 						{
// 							ID:   "ID2",
// 							Type: 1,
// 						},
// 					},
// 				},
// 			},
// 			result: "U2",
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			sMock := m.NewMockdiscordSession(ctrl)
// 			c := Client{targetRoleID: tt.tr}

// 			sMock.
// 				EXPECT().
// 				Channel(tt.discordContract.ic.ChannelID).
// 				Times(1).
// 				Return(tt.discordContract.ch, tt.discordContract.chErr)

// 			sMock.
// 				EXPECT().
// 				GuildMembers(tt.discordContract.guildID, "", 1000).
// 				Times(1).
// 				Return(tt.discordContract.members, tt.discordContract.mErr)

// 			ir := &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Content: tt.result,
// 				},
// 			}

// 			sMock.
// 				EXPECT().
// 				InteractionRespond(tt.discordContract.ic.Interaction, ir).
// 				Times(1).
// 				Return(tt.discordContract.iErr)

// 			c.shuffleChannelMembers(sMock, tt.discordContract.ic)
// 		})
// 	}
// }
