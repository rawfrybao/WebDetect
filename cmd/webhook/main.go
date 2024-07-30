package main

import (
	"log"
	"os"
	"webdetect/internal/logger"
	"webdetect/internal/server"
	"webdetect/internal/webhook"
)

func main() {
	err := webhook.SetupWebhook()
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())
		os.Exit(1)
	}

	go webhook.Monitor()

	server.Start()
}
