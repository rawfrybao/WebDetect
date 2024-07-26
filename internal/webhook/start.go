package webhook

import (
	"fmt"
	"webdetect/internal/api"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	fmt.Println("Start command", *req.Message.From.ID)
}
