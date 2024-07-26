package webhook

import (
	"fmt"
	"time"
	"webdetect/internal/db"
	"webdetect/internal/detect"
)

func Monitor() {
	for {
		tasks, err := db.GetTasks()
		if err != nil {
			fmt.Println(err)
		}

		for _, task := range tasks {
			content := detect.GetContent(task.URL, task.XPath)

			if content == task.PrevContent {
				continue
			}

			prev_content := task.PrevContent

			err = db.SetPrevContent(task.ID, content)
			if err != nil {
				fmt.Println(err)
			}
			go NotifyUsers(task.ID, content, prev_content)
		}

		time.Sleep(5 * time.Minute)
	}
}
