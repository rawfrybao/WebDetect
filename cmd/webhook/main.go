package main

import (
	"fmt"
	"os"
	"webdetect/internal/server"
)

func main() {
	tgBotId := os.Getenv("TG_BOT_TOKEN")
	if tgBotId == "" {
		fmt.Println("TG_BOT_TOKEN is not set")
		os.Exit(1)
	}

	apiURL := os.Getenv("TG_API_URL")
	if apiURL == "" {
		fmt.Println("TG_API_URL is not set")
		os.Exit(1)
	}

	err := webhook.SetupWebhook()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go webhook.Monitor()

	server.Start()
}
