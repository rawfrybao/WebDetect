package db

type Task struct {
	ID          int64
	URL         string
	XPath       string
	PrevContent string
}

type User struct {
	ID        int64
	TGID      int64
	ChatID    int64
	IsAdmin   bool
	HasAccess bool
}

type Subscription struct {
	ID     int64
	Name   string
	UserID int64
	TaskID int64
}
