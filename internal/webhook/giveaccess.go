package webhook

import (
	"log"
	"strconv"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func handleGiveAccess(req api.HandleUpdateJSONRequestBody, text string) {
	isAdmin, err := db.CheckIsAdmin(*req.Message.From.ID)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())
		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "遇到bug了"

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	if !isAdmin {
		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "你不行"

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ReplyParameters: &replyParams,
		}

		go SendMessage(message)
		return
	}

	if text == "" && req.Message.ReplyToMessage == nil {
		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "回复一个消息"

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ReplyParameters: &replyParams,
		}

		go SendMessage(message)
		return
	}

	var targetId int64

	if text != "" {
		idInt, err := strconv.Atoi(text)
		if err != nil {
			log.Println(err.Error())
			logger.Log(err.Error())

			replyParams := TgMessageReplyParameters{
				MessageID: *req.Message.MessageID,
				ChatID:    *req.Message.Chat.ID,
			}

			msg := "回复一个消息或者输入对方的TG ID"

			message := TgMessage{
				ChatID:          *req.Message.Chat.ID,
				Text:            msg,
				ReplyParameters: &replyParams,
			}

			go SendMessage(message)
			return
		}

		targetId = int64(idInt)
	} else {
		targetId = *req.Message.ReplyToMessage.From.ID
	}

	err = db.SetHasAccess(targetId, true)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "遇到bug了"

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := "给权限成功"

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            msg,
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
