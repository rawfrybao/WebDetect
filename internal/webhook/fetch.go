package webhook

import (
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/detect"
	"webdetect/internal/logger"
)

func handleFetch(req api.HandleUpdateJSONRequestBody, text string) {
	if len(strings.Split(text, " ")) != 2 {
		logger.Log("Invalid number of arguments")
	}

	content, err := detect.Fetch(text)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(content)

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            content,
		ParseMode:       "MarkdownV2",
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
