package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/psy1992/shuffle-chat/internal/discord"
)

const runCommand = "Please run: env DISCORD_TOKEN=<bot token> TARGET_ROLE_ID=<role id> ./app"

var token string
var targetRoleID string

func init() {
	token = os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("No token provided. " + runCommand)
	}

	targetRoleID = os.Getenv("TARGET_ROLE_ID")
	if targetRoleID == "" {
		log.Fatal("No targetRoleID provided. " + runCommand)
	}
}

func main() {
	discordClient := discord.NewClient(token, targetRoleID)
	defer discordClient.Session.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Shuffler is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
