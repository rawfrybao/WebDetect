package main

import (
	"os"
	"webdetect/internal/logger"
	"webdetect/internal/server"
	"webdetect/internal/webhook"
)

func main() {
	err := webhook.SetupWebhook()
	if err != nil {
		logger.Log(err.Error())
		os.Exit(1)
	}

	go webhook.Monitor()

	server.Start()
}
