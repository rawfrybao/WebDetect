package webhook

import (
	"webdetect/internal/api"
	"webdetect/internal/logger"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	logger.Log("Start command", *req.Message.From.ID)
}
