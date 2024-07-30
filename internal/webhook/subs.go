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
			ChatID: *req.Message.Chat.ID,
			Text:   msg,

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
			ChatID: *req.Message.Chat.ID,
			Text:   msg,

			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	if len(strings.Split(text, " ")) != 3 {
		log.Println("Invalid number of arguments")
		logger.Log("Invalid number of arguments")
		msg := "使用方法  \n"
		msg += "/addsub <name> <url> <xpath>  \n"

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

		msg := "添加失败  \n"
		msg += err.Error()

		message := TgMessage{
			ChatID: *req.Message.Chat.ID,
			Text:   msg,

			ReplyParameters: &replyParams,
		}
		go SendMessage(message)
		return
	}

	log.Println("Subscription added", name, url, xpath)
	logger.Log("Subscription added", name, url, xpath)

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := "添加 " + name + " 成功"

	message := TgMessage{
		ChatID: *req.Message.Chat.ID,
		Text:   msg,

		ReplyParameters: &replyParams,
	}
	go SendMessage(message)
}

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
			ChatID: *req.Message.Chat.ID,
			Text:   msg,

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
		ChatID: *req.Message.Chat.ID,
		Text:   msg,

		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}

func handleListSubscription(req api.HandleUpdateJSONRequestBody) {
	tgID := *req.Message.From.ID

	subscriptions, err := db.GetSubscriptions(tgID)
	if err != nil {
		log.Println(err.Error())
		logger.Log(err.Error())

		replyParams := TgMessageReplyParameters{
			MessageID: *req.Message.MessageID,
			ChatID:    *req.Message.Chat.ID,
		}

		msg := "获取失败  \n"
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

	msg := "订阅列表  \n"
	for _, sub := range subscriptions {
		msg += sub + "  \n"
	}

	message := TgMessage{
		ChatID:          *req.Message.Chat.ID,
		Text:            msg,
		ReplyParameters: &replyParams,
	}

	go SendMessage(message)
}
