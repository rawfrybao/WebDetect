package webhook

import (
	"log"
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/detect"
	"webdetect/internal/logger"
)

func handleFetch(req api.HandleUpdateJSONRequestBody, text string) {
	if len(strings.Split(text, " ")) != 2 {
		log.Println("Invalid number of arguments")
		logger.Log("Invalid number of arguments")

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "参数不对  \n"
		msg += "使用方法：  \n"
		msg += "/fetch <url> <xpath>  \n"

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ReplyParameters: &replyParams,
		}

		go SendMessage(message)
		return
	}

	content, err := detect.Fetch(text)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		msg := "获取失败  \n"
		msg += err.Error()

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		message := TgMessage{
			ChatID: *req.Message.Chat.ID,
			Text:   msg,

			ReplyParameters: &replyParams,
		}

		go SendMessage(message)
		return
	}
	log.Println(content)
	logger.Log(content)

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	message := TgMessage{
		ChatID: *req.Message.Chat.ID,
		Text:   content,

		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
