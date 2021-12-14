package mydomain

import "fmt"

type Todo struct {
	Title     string
	Completed bool
}

func NewTodo(title string) *Todo {
	fmt.Printf("@ domain: new todo %q\n", title)
	return &Todo{Title: title}
}

func (todo Todo) String() string {
	return todo.Title
}
