package db

import (
	"context"
	"fmt"
)

func NewTask(url string, xpath string) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if task exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM tasks WHERE url = $1 AND xpath = $2)", url, xpath).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if exists {
		return nil
	}

	prev_content := ""
	_, err = conn.Exec(context.Background(), "INSERT INTO tasks (url, xpath, prev_content) VALUES ($1, $2, $3)", url, xpath, prev_content)
	if err != nil {
		return fmt.Errorf("could not insert into database: %w", err)
	}

	return nil
}

func GetTasks() ([]Task, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("could not query database: %w", err)
	}

	var tasks []Task
	for rows.Next() {
		var id int64
		var url string
		var xpath string
		var prev string
		err = rows.Scan(&id, &url, &xpath, &prev)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}

		tasks = append(tasks, Task{
			ID:          id,
			URL:         url,
			XPath:       xpath,
			PrevContent: prev,
		})
	}

	return tasks, nil
}

func GetTask(taskId int64) (Task, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return Task{}, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	var task Task
	err = conn.QueryRow(context.Background(), "SELECT * FROM tasks WHERE id = $1", taskId).Scan(&task.ID, &task.URL, &task.XPath, &task.PrevContent)
	if err != nil {
		return Task{}, fmt.Errorf("could not query database: %w", err)
	}

	return task, nil
}

func CheckTaskHasSubscriptions(taskId int64) (bool, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return false, fmt.Errorf("could not connect to database: %w", err)
	}

	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM subscriptions WHERE task_id = $1)", taskId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not query database: %w", err)
	}

	return exists, nil
}

func DeleteTask(taskId int64) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", taskId)
	if err != nil {
		return fmt.Errorf("could not delete from database: %w", err)
	}

	return nil
}

func CheckTaskExists(url, xpath string) (bool, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return false, fmt.Errorf("could not connect to database: %w", err)
	}

	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM tasks WHERE url = $1 AND xpath = $2)", url, xpath).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not query database: %w", err)
	}

	return exists, nil
}

func GetTaskByInfo(url, xpath string) (int64, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return 0, fmt.Errorf("could not connect to database: %w", err)
	}

	var id int64
	err = conn.QueryRow(context.Background(), "SELECT id FROM tasks WHERE url = $1 AND xpath = $2", url, xpath).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not query database: %w", err)
	}

	return id, nil
}

func SetPrevContent(taskId int64, content string) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	_, err = conn.Exec(context.Background(), "UPDATE tasks SET prev_content = $1 WHERE id = $2", content, taskId)
	if err != nil {
		return fmt.Errorf("could not update database: %w", err)
	}

	return nil
}
