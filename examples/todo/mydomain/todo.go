package mydomain

import "fmt"

type Todo struct {
	Title       string
	isCompleted bool
}

func NewTodo(title string) *Todo {
	fmt.Printf("@ domain: new todo %q\n", title)
	return &Todo{Title: title}
}

func (todo Todo) String() string {
	return todo.Title
}

func (todo Todo) IsDone() bool {
	return todo.isCompleted
}
