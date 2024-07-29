package webhook

import (
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func NotifyUsers(taskId int64, content, prev_content string) {
	users, err := db.GetUsersWithTask(taskId)
	if err != nil {
		logger.Log(err)
		return
	}

	for _, user := range users {
		if user.ChatID == -1 {
			continue
		}

		subscription, err := db.GetSubscriptionByTaskID(user.ID, taskId)
		if err != nil {
			logger.Log(err)
			continue
		}

		task, err := db.GetTask(taskId)
		if err != nil {
			logger.Log(err)
			continue
		}

		logger.Log(subscription.Name, prev_content, content)
		// Notify user
		text := "#" + subscription.Name + "  \nFrom: " + prev_content + "  \nTo: " + content + "  \n[Link](" + task.URL + ")"

		message := TgMessage{
			ChatID:    user.ChatID,
			Text:      text,
			ParseMode: "MarkdownV2",
		}

		go SendMessage(message)
	}
}
