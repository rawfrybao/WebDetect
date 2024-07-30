package webhook

import (
	"log"
	"strconv"
	"webdetect/internal/api"
	"webdetect/internal/logger"
)

func handleMyID(req api.HandleUpdateJSONRequestBody) {
	log.Println(*req.Message.From.ID, "Requesting ID")
	logger.Log(*req.Message.From.ID, "Requesting ID")

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := strconv.Itoa(int(*req.Message.From.ID))

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            msg,
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
