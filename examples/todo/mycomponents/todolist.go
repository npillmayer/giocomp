package mycomponents

import (
	"fmt"

	"gioui.org/io/event"
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
	list          *mydomain.TodoList
	todoDelegates []*TodoDelegate
	state         *widget.List
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
	case components.UpdateEvent:
		if l.list == nil {
			return
		}
		if e.Target == l {
			l.Async(nil, func() interface{} {
				fmt.Printf("@ todo list event %#v\n", ev)
				todo := e.Value.(*mydomain.Todo)
				l.list.AppendTodo(todo)
				l.AdaptTodoDelegatesFromList()
				fmt.Printf("@ now del-list = %v\n", l.todoDelegates)
				return nil
			})
		}
	}
}

func (l *TodoListDelegate) AdaptTodoDelegatesFromList() {
	todos := l.list.Todos()
	if len(l.todoDelegates) < len(todos) {
		j := len(todos) - len(l.todoDelegates)
		l.todoDelegates = append(l.todoDelegates, make([]*TodoDelegate, j)...)
	} else if len(l.todoDelegates) > len(todos) {
		l.todoDelegates = l.todoDelegates[:len(todos)]
	}
	for i := range l.todoDelegates {
		l.todoDelegates[i] = NewTodoDelegate(todos[i])
	}
}

// --- Rendering -------------------------------------------------------------

type (
	D = layout.Dimensions
	C = layout.Context
)

func TodoList(l *TodoListDelegate) layout.Widget {
	todos := make([]layout.Widget, len(l.todoDelegates))
	for i, todo := range l.todoDelegates {
		todos[i] = Todo(todo)
	}
	return func(gtx C) D {
		return material.List(html.Theme.Material(), l.state).Layout(gtx, len(todos), func(gtx C, i int) D {
			return todos[i](gtx)
		})
	}
}

/*
func H2Todo(t *mydomain.Todo) layout.Widget {
	return html.H2().Text(t.Title)
}
*/
