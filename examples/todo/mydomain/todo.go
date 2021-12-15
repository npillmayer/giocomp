package mydomain

type Todo struct {
	Title     string
	Completed bool
}

func NewTodo(title string) *Todo {
	return &Todo{Title: title}
}

func (todo Todo) String() string {
	return todo.Title
}
