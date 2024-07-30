package webhook

import (
	"log"
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func handleDeleteSubscription(req api.HandleUpdateJSONRequestBody, text string) {
	hasAccess, err := db.CheckHasAccess(*req.Message.From.ID)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())
		msg := "遇到bug了"

		message := TgMessage{
			ChatID:    *req.Message.Chat.ID,
			Text:      msg,
			ParseMode: "MarkdownV2",
		}
		go SendMessage(message)
		return
	}

	if !hasAccess {
		log.Println("User does not have access")
		logger.Log("User does not have access")
		msg := "你不行"

		message := TgMessage{
			ChatID:    *req.Message.Chat.ID,
			Text:      msg,
			ParseMode: "MarkdownV2",
		}
		go SendMessage(message)
		return
	}

	name := strings.Split(text, " ")[0]

	err = db.DeleteSubscription(*req.Message.From.ID, name)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "删除失败  \n"
		msg += err.Error()

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ParseMode:       "MarkdownV2",
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := "删除成功（也许）"

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            msg,
		ParseMode:       "MarkdownV2",
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
