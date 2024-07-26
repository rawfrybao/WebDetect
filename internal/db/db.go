package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return conn, nil
}

func GetUserIDByTGID(tgId int64) (int64, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return 0, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	var userId int64
	err = conn.QueryRow(context.Background(), "SELECT id FROM users WHERE tg_id = $1", tgId).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("could not query database: %w", err)
	}

	return userId, nil
}

func GetTGIDByUserID(userId int64) (int64, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return 0, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	var tgId int64
	err = conn.QueryRow(context.Background(), "SELECT tg_id FROM users WHERE id = $1", userId).Scan(&tgId)
	if err != nil {
		return 0, fmt.Errorf("could not query database: %w", err)
	}

	return tgId, nil
}

func GetChatIDByUserID(userId int64) (int64, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return 0, fmt.Errorf("could not connect to database: %w", err)
	}

	var chatId int64
	err = conn.QueryRow(context.Background(), "SELECT chat_id FROM users WHERE id = $1", userId).Scan(&chatId)
	if err != nil {
		return 0, fmt.Errorf("could not query database: %w", err)
	}

	return chatId, nil
}

func GetSubscriptionsByUser(userId int64) ([]int64, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT task_id FROM subscriptions WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("could not query database: %w", err)
	}
	defer rows.Close()

	var subscriptions []int64
	for rows.Next() {
		var taskId int64
		err = rows.Scan(&taskId)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		subscriptions = append(subscriptions, taskId)
	}

	return subscriptions, nil
}

func GetUsersWithTaskSubscription(taskId int64) ([]string, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	rows, err := conn.Query(context.Background(), "SELECT user_id FROM subscriptions WHERE task_id = $1", taskId)
	if err != nil {
		return nil, fmt.Errorf("could not query database: %w", err)
	}

	var userIds []string
	for rows.Next() {
		var userId string
		err = rows.Scan(&userId)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		userIds = append(userIds, userId)
	}

	return userIds, nil
}
