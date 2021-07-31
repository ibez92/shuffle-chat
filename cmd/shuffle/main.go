package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"shufflezoommeeting/internal/discord"
	"syscall"
)

const runCommand = "Please run: env DISCORD_TOKEN=<bot token> GUILD_ID=<guild id> ./app"

var token string
var guildID string

func init() {
	token = os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("No token provided. " + runCommand)
	}

	guildID = os.Getenv("GUILD_ID")
	if guildID == "" {
		log.Fatal("No guildID provided. " + runCommand)
	}
}

func main() {
	discordClient := discord.NewClient(token, guildID)
	defer discordClient.Session.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Shuffler is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
