package webhook

import (
	"log"
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func addSubscription(tgID int64, name, url, xpath string) error {

	err := db.NewSubscription(tgID, name, url, xpath)
	if err != nil {
		return err
	}

	return nil
}

func handleAddSubscription(req api.HandleUpdateJSONRequestBody, text string) {
	hasAccess, err := db.CheckHasAccess(*req.Message.From.ID)
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
			ParseMode:       "MarkdownV2",
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	if !hasAccess {
		log.Println("User does not have access")
		logger.Log("User does not have access")
		msg := "你不行  \n"

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ParseMode:       "MarkdownV2",
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	if len(strings.Split(text, " ")) != 3 {
		log.Println("Invalid number of arguments")
		logger.Log("Invalid number of arguments")
	}

	args := strings.Split(text, " ")
	name := args[0]
	url := args[1]
	xpath := args[2]

	err = addSubscription(*req.Message.From.ID, name, url, xpath)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := err.Error()

		message := TgMessage{
			ChatID:          *req.Message.Chat.ID,
			Text:            msg,
			ParseMode:       "MarkdownV2",
			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	log.Println("Subscription added", name, url, xpath)
	logger.Log("Subscription added", name, url, xpath)

	msg := "添加 " + name + " 成功"

	message := TgMessage{
		ChatID:    *req.Message.Chat.ID,
		Text:      msg,
		ParseMode: "MarkdownV2",
	}
	go SendMessage(message)

}
