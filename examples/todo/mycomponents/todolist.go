package mycomponents

import (
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

type TodoListDelegate struct {
	components.Worker
	filterCompleted bool
	list            *mydomain.TodoList
	todoDelegates   map[*mydomain.Todo]*TodoDelegate
	state           *widget.List
}

func NewTodoListDelegate(list *mydomain.TodoList) *TodoListDelegate {
	return &TodoListDelegate{
		list: list,
		state: &widget.List{
			List: layout.List{Axis: layout.Vertical},
		},
	}
}

func (l *TodoListDelegate) Event(ev event.Event) {
	switch e := ev.(type) {
	case giocomp.ConnectEvent:
		l.Connect(e.Connection())
		return
	case system.FrameEvent:
		for _, t := range l.todoDelegates {
			t.Event(ev)
		}
	case components.UpdateEvent:
		if l.list == nil {
			return
		}
		if e.Target == l {
			l.Async(nil, func() interface{} {
				todo := e.Value.(*mydomain.Todo)
				l.list.AppendTodo(todo)
				l.AdaptTodoDelegatesFromList()
				return nil
			})
		}
	}
}

func (l *TodoListDelegate) AdaptTodoDelegatesFromList() {
	if l.todoDelegates == nil {
		l.todoDelegates = make(map[*mydomain.Todo]*TodoDelegate)
	}
	for _, t := range l.list.Todos() {
		if _, ok := l.todoDelegates[t]; !ok {
			l.todoDelegates[t] = NewTodoDelegate(t)
		}
	}
}

func (l *TodoListDelegate) FilterCompleted(doFilter bool) {
	l.filterCompleted = doFilter
	l.Async(nil, func() interface{} {
		return nil
	})
}

// --- Rendering -------------------------------------------------------------

func TodoList(l *TodoListDelegate) layout.Widget {
	count := 0
	todos := make([]layout.Widget, len(l.todoDelegates))
	for _, todo := range l.list.Todos() {
		if l.filterCompleted && todo.Completed {
			continue
		}
		todos[count] = Todo(l.todoDelegates[todo])
		count++
	}
	return func(gtx layout.Context) layout.Dimensions {
		return material.List(html.Theme.Material(), l.state).Layout(gtx, count,
			func(gtx layout.Context, i int) layout.Dimensions {
				return todos[i](gtx)
			})
	}
}
