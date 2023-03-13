package helper

import "time"

type Todo struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"due_date"`
	IsComplete  bool      `json:"is_complete"`
	UserId      uint      `json:"user_id"`
	CreateAt    time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

type Message struct {
	Todo      Todo   `json:"todo"`
	UserEmail string `json:"user_email"`
	Username  string `json:"username"`
}
