package webhook

import (
	"fmt"
	"webdetect/internal/api"
	"webdetect/internal/logger"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	fmt.Println("Start command", *req.Message.From.ID)
	logger.Log("Start command", *req.Message.From.ID)
}
