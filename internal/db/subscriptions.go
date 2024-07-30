package db

import (
	"context"
	"fmt"
)

func NewSubscription(tgID int64, name string, url string, xpath string) error {

	userID, err := GetUserIDByTGID(tgID)
	if err != nil {
		return fmt.Errorf("could not get user id by tg id: %w", err)
	}

	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if subscription exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM subscriptions WHERE name = $1)", name).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if exists {
		return fmt.Errorf("subscription with name %s already exists", name)
	}

	exists, err = CheckTaskExists(url, xpath)
	if err != nil {
		return fmt.Errorf("could not check task exists: %w", err)
	}

	if !exists {
		err = NewTask(url, xpath)
		if err != nil {
			return fmt.Errorf("could not create new task: %w", err)
		}
	}

	taskId, err := GetTaskByInfo(url, xpath)
	if err != nil {
		return fmt.Errorf("could not get task by info: %w", err)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO subscriptions (name, user_id, task_id) VALUES ($1, $2, $3)", name, userID, taskId)
	if err != nil {
		return fmt.Errorf("could not insert into database: %w", err)
	}

	return nil
}

func DeleteSubscription(tgID int64, name string) error {

	userID, err := GetUserIDByTGID(tgID)
	if err != nil {
		return fmt.Errorf("could not get user id by tg id: %w", err)
	}

	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if subscription exists, and get task id
	var taskID int64
	err = conn.QueryRow(context.Background(), "SELECT task_id FROM subscriptions WHERE name = $1 AND user_id = $2", name, userID).Scan(&taskID)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	_, err = conn.Exec(context.Background(), "DELETE FROM subscriptions WHERE name = $1 AND user_id = $2", name, userID)
	if err != nil {
		return fmt.Errorf("could not delete from database: %w", err)
	}

	// Check if task has other subscriptions
	hasSubs, err := CheckTaskHasSubscriptions(taskID)
	if err != nil {
		return fmt.Errorf("could not check task has subscriptions: %w", err)
	}

	if !hasSubs {
		err = DeleteTask(taskID)
		if err != nil {
			return fmt.Errorf("could not delete task: %w", err)
		}
	}

	return nil
}
