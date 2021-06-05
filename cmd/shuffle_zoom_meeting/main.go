package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"shufflezoommeeting/internal/discord"
	"syscall"
)

var token string
var zoomToken string
var zoomSecret string
var zoomMeetingID string

func init() {
	token = os.Getenv("DISCORD_TOKEN")
	zoomToken = os.Getenv("ZOOM_TOKEN")
	zoomSecret = os.Getenv("ZOOM_SECRET")
	zoomMeetingID = os.Getenv("ZOOM_MEETING_ID")
}

func main() {
	validateTokens()

	dg, err := discord.Open(token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}
	// Cleanly close down the Discord session.
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("shuffle-zoom-conf is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func validateTokens() {
	if token == "" {
		log.Fatal("No token provided. Please run: shuffle-zoom-conf -t <bot token>")
	}

	if zoomToken == "" {
		log.Fatal("No zoom token provided. Please run: shuffle-zoom-conf -zoomt <zoom token>")
	}

	if zoomSecret == "" {
		log.Fatal("No token provided. Please run: shuffle-zoom-conf -zooms <zoom secret>")
	}
}
