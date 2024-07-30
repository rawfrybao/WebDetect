package db

import (
	"context"
	"fmt"
)

func NewUser(tgId, chatId int64) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if user exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)", tgId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if exists {
		return fmt.Errorf("user with tg_id %d already exists", tgId)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO users (tg_id, chat_id) VALUES ($1, $2)", tgId, chatId)
	if err != nil {
		return fmt.Errorf("could not insert into database: %w", err)
	}

	return nil
}

func GetUsers() ([]User, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("could not query database: %w", err)
	}

	var users []User
	for rows.Next() {
		var id int64
		var tgId int64
		var isAdmin bool
		var hasAccess bool
		err = rows.Scan(&id, &tgId, &isAdmin, &hasAccess)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}

		users = append(users, User{
			ID:        id,
			TGID:      tgId,
			IsAdmin:   isAdmin,
			HasAccess: hasAccess,
		})
	}

	return users, nil
}

func SetIsAdmin(tgId int64, isAdmin bool) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if user exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)", tgId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if !exists {
		_, err = conn.Exec(context.Background(), "INSERT INTO users (tg_id, is_admin) VALUES ($1, $2)", tgId, isAdmin)
		if err != nil {
			return fmt.Errorf("could not insert into database: %w", err)
		}
		return nil
	}

	_, err = conn.Exec(context.Background(), "UPDATE users SET is_admin = $1 WHERE tg_id = $2", isAdmin, tgId)
	if err != nil {
		return fmt.Errorf("could not update database: %w", err)
	}

	return nil
}

func SetHasAccess(tgId int64, hasAccess bool) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if user exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)", tgId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if !exists {
		_, err = conn.Exec(context.Background(), "INSERT INTO users (tg_id, has_access) VALUES ($1, $2)", tgId, hasAccess)
		if err != nil {
			return fmt.Errorf("could not insert into database: %w", err)
		}
		return nil
	}

	_, err = conn.Exec(context.Background(), "UPDATE users SET has_access = $1 WHERE tg_id = $2", hasAccess, tgId)
	if err != nil {
		return fmt.Errorf("could not update database: %w", err)
	}

	return nil
}

func SetChat(tgId int64, chatId int64) error {
	conn, err := NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	// Check if user exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)", tgId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not query database: %w", err)
	}

	if !exists {
		_, err = conn.Exec(context.Background(), "INSERT INTO users (tg_id, chat_id) VALUES ($1, $2)", tgId, chatId)
		if err != nil {
			return fmt.Errorf("could not insert into database: %w", err)
		}
		return nil
	}

	_, err = conn.Exec(context.Background(), "UPDATE users SET chat_id = $1 WHERE tg_id = $2", chatId, tgId)
	if err != nil {
		return fmt.Errorf("could not update database: %w", err)
	}

	return nil
}

func CheckIsAdmin(tgId int64) (bool, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return false, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	var isAdmin bool
	err = conn.QueryRow(context.Background(), "SELECT is_admin FROM users WHERE tg_id = $1", tgId).Scan(&isAdmin)
	if err != nil {
		return false, fmt.Errorf("could not query database: %w", err)
	}

	return isAdmin, nil
}

func CheckHasAccess(tgId int64) (bool, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return false, fmt.Errorf("could not connect to database: %w", err)
	}

	var hasAccess bool
	err = conn.QueryRow(context.Background(), "SELECT has_access FROM users WHERE tg_id = $1", tgId).Scan(&hasAccess)
	if err != nil {
		return false, fmt.Errorf("could not query database: %w", err)
	}

	return hasAccess, nil
}

func GetUsersWithTask(taskId int64) ([]User, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM users WHERE id IN (SELECT user_id FROM subscriptions WHERE task_id = $1)", taskId)
	if err != nil {
		return nil, fmt.Errorf("could not query database: %w", err)
	}

	var users []User
	for rows.Next() {
		var id int64
		var tgId int64
		var chatId int64
		var isAdmin bool
		var hasAccess bool
		err = rows.Scan(&id, &tgId, &chatId, &isAdmin, &hasAccess)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}

		users = append(users, User{
			ID:        id,
			TGID:      tgId,
			ChatID:    chatId,
			IsAdmin:   isAdmin,
			HasAccess: hasAccess,
		})
	}

	return users, nil
}

func GetSubscriptionByTaskID(userId, taskId int64) (Subscription, error) {
	conn, err := NewPostgresConnection()
	if err != nil {
		return Subscription{}, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	var subscription Subscription
	err = conn.QueryRow(context.Background(), "SELECT * FROM subscriptions WHERE user_id = $1 AND task_id = $2", userId, taskId).Scan(&subscription.ID, &subscription.Name, &subscription.UserID, &subscription.TaskID)
	if err != nil {
		return Subscription{}, fmt.Errorf("could not query database: %w", err)
	}

	return subscription, nil
}
