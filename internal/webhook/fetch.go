package webhook

import (
	"fmt"
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/detect"
)

func handleFetch(req api.HandleUpdateJSONRequestBody, text string) {
	if len(strings.Split(text, " ")) != 2 {
		fmt.Println("Invalid number of arguments")
	}

	content, err := detect.Fetch(text)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(content)

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
