package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot token")
	flag.Parse()
}

var token string

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: sortmyvoice -t <bot token>")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	// Cleanly close down the Discord session.
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Sortmyvoice is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "!sortmyvoice")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "!sortmyvoice" {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	if c.Name != "bot" {
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return
	}

	var authorVoiceChID string
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			authorVoiceChID = vs.ChannelID
		}
	}

	userNames := make([]string, 15)
	counter := 0
	for _, vs := range g.VoiceStates {
		if vs.ChannelID == authorVoiceChID {
			m, err := s.State.Member(vs.GuildID, vs.UserID)
			if err != nil {
				continue
			}
			userNames[counter] = fmt.Sprintf("%d) %s", counter+1, m.User.String())
			counter++
		}
	}

	sort.Strings(userNames)
	s.ChannelMessageSend(c.ID, strings.Join(userNames, "\n"))
}
