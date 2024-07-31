package webhook

import (
	"log"
	"strconv"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func handleInfo(req api.HandleUpdateJSONRequestBody) {
	log.Println("Info command", *req.Message.From.ID)
	logger.Log("Info command", *req.Message.From.ID)

	// Check if the user is in the database
	user, err := db.GetUserByTGID(*req.Message.From.ID)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "遇到bug了  \n"
		msg += err.Error()

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

	msg := "用户信息  \n"
	msg += "ID: " + strconv.FormatInt(user.ID, 10) + "  \n"
	msg += "TGID: " + strconv.FormatInt(user.TGID, 10) + "  \n"
	msg += "ChatID: " + strconv.FormatInt(user.ChatID, 10) + "  \n"
	msg += "IsAdmin: " + strconv.FormatBool(user.IsAdmin) + "  \n"
	msg += "HasAccess: " + strconv.FormatBool(user.HasAccess) + "  \n"

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            msg,
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
