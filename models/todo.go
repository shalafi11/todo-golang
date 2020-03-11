package models

import (
	"time"
)

var todoLastID int = 0

// Todo simple todo model
type Todo struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
}

// TodoCollection collects all todoitems
type TodoCollection struct {
	TodoCollection []Todo
}

// NewTodo is a simple Todo constructor
func NewTodo(title, body string) *Todo {
	todo := new(Todo)
	todo.ID = todoLastID
	todo.Title = title
	todo.Body = body
	todo.CreatedAt = time.Now()

	todoLastID++
	return todo
}

// NewTodoCollection is a simple TodoCollection constructor
func NewTodoCollection() *TodoCollection {
	return &TodoCollection{}
}

//Add adding new intems to TodoCollection
func (t *TodoCollection) Add(todo *Todo) bool {
	temp := len(t.TodoCollection)
	t.TodoCollection = append(t.TodoCollection, *todo)

	if len(t.TodoCollection) == temp+1 {
		return true
	}
	return false
}

// Retrieve one element from TodoCollection with given id param
func (t *TodoCollection) Retrieve(id int) *Todo {
	for _, todo := range t.TodoCollection {
		if id == todo.ID {
			return &todo
		}
	}
	return nil
}

// RetrieveAll elements from TodoCollection
func (t *TodoCollection) RetrieveAll() []Todo {
	return t.TodoCollection
}

// Update todo with given id
func (t *TodoCollection) Update(id int, title, body string) bool {
	for i, todo := range t.TodoCollection {
		if id == todo.ID {
			(&t.TodoCollection[i]).Title = title
			(&t.TodoCollection[i]).Body = body
			return true
		}
	}
	return false
}

// Delete element from TodoCollection
func (t *TodoCollection) Delete(id int) bool {
	var index int

	for i, todo := range t.TodoCollection {
		if id == todo.ID {
			index = i
			break
		}
		return false
	}
	temp := len(t.TodoCollection)
	t.TodoCollection = t.TodoCollection[:index+copy(t.TodoCollection[index:], t.TodoCollection[index+1:])]
	if temp-1 == len(t.TodoCollection) {
		return true
	}
	return false

}
