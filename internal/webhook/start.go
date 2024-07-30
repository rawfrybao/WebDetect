package webhook

import (
	"log"
	"webdetect/internal/api"
	"webdetect/internal/logger"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	log.Println("Start command", *req.Message.From.ID)
	logger.Log("Start command", *req.Message.From.ID)
}
