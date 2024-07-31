package webhook

import (
	"log"
	"strconv"
	"time"
	"webdetect/internal/db"
	"webdetect/internal/detect"
	"webdetect/internal/logger"
)

func Monitor() {
	for {
		tasks, err := db.GetTasks()
		if err != nil {
			log.Println(err)
			logger.Log(err)

			time.Sleep(5 * time.Minute)
			continue
		}

		for _, task := range tasks {
			content := detect.GetContent(task.URL, task.XPath)

			logger.Log("Task", strconv.FormatInt(task.ID, 10), "Content", content, "PrevContent", task.PrevContent)

			if content == task.PrevContent {
				logger.Log(content, "is the same as", task.PrevContent)
				continue
			}

			prev_content := task.PrevContent

			err = db.SetPrevContent(task.ID, content)
			if err != nil {
				log.Println(err)
				logger.Log(err)
			}
			go NotifyUsers(task.ID, content, prev_content)
		}

		time.Sleep(5 * time.Minute)
	}
}
