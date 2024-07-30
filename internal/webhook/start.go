package webhook

import (
	"log"
	"webdetect/internal/api"
	"webdetect/internal/logger"
)

func handleStart(req api.HandleUpdateJSONRequestBody) {
	log.Println("Start command", *req.Message.From.ID)
	logger.Log("Start command", *req.Message.From.ID)

	replyParams := TgMessageReplyParameters{
		MessageID: *req.Message.MessageID,
		ChatID:    *req.Message.Chat.ID,
	}

	msg := "这玩意你都用？  \n"
	msg += "你可以使用以下命令：  \n"
	msg += "/fetch <url> <xpath> - 获取网页内容  \n"
	msg += "/listsub - 列出所有订阅  \n"
	msg += "/addsub <name> <url> <xpath> - 添加订阅  \n"
	msg += "/delsub <name> - 删除订阅  \n"

	message := TgMessage{
		ChatID: *req.Message.Chat.ID,
		Text:   msg,

		ReplyParameters: &replyParams,
	}

	SendMessage(message)

}
