package models

import "time"

// Todo simple todo model
type Todo struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

// NewTodo is a simple Todo constructor
func NewTodo(title string, body string) *Todo {
	todo := new(Todo)
	todo.Title = title
	todo.Body = body
	todo.CreatedAt = time.Now()

	return todo
}
