package webhook

import (
	"log"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func NotifyUsers(taskId int64, content, prev_content string) {
	users, err := db.GetUsersWithTask(taskId)
	if err != nil {
		log.Println(err)
		logger.Log(err)
		return
	}

	for _, user := range users {
		if user.ChatID == -1 {
			continue
		}

		subscription, err := db.GetSubscriptionByTaskID(user.ID, taskId)
		if err != nil {
			log.Println(err)
			logger.Log(err)
			continue
		}

		task, err := db.GetTask(taskId)
		if err != nil {
			log.Println(err)
			logger.Log(err)
			continue
		}

		log.Println("Notifying", subscription.Name, prev_content, content)
		logger.Log("Notifying", subscription.Name, prev_content, content)
		// Notify user
		msg := subscription.Name + "  \n"
		msg += "From: " + prev_content + "  \n"
		msg += "To: " + content + "  \n"
		msg += "Link: " + task.URL + "  \n"

		message := TgMessage{
			ChatID: user.ChatID,
			Text:   msg,
		}

		go SendMessage(message)
	}
}
