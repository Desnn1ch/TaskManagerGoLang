package models

type Task struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Status string `db:"status"`
	UserID int    `db:"user_id"`
}
