package webhook

import (
	"log"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	log.Println("Start command", *req.Message.From.ID)
	logger.Log("Start command", *req.Message.From.ID)

	// Check if the user is in the database
	_, err := db.GetUserIDByTGID(*req.Message.From.ID)
	if err != nil {
		// If the user is not in the database, add them
		err := db.NewUser(*req.Message.From.ID, *req.Message.Chat.ID)
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
	}

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := "这玩意你都用？  \n"
	msg += "命令列表：  \n"
	msg += "/fetch <url> <xpath> - 获取网页内容  \n"
	msg += "/listsub - 列出所有订阅  \n"
	msg += "/addsub <name> <url> <xpath> - 添加订阅  \n"
	msg += "/delsub <name> - 删除订阅  \n"

	message := TgMessage{
		ChatID: *req.Message.Chat.ID,
		Text:   msg,

		ReplyParameters: &replyParams,
	}

	go SendMessage(message)

}
