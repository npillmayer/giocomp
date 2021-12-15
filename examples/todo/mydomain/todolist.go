package mydomain

type TodoList struct {
	todos []*Todo
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos: make([]*Todo, 0, 10),
	}
}

func (l TodoList) Len() int {
	return len(l.todos)
}

func (l TodoList) Todos() []*Todo {
	return l.todos
}

func (l *TodoList) AppendTodo(todo *Todo) {
	l.todos = append(l.todos, todo)
}
