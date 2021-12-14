package mycomponents

import (
	"fmt"

	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

type TodoDelegate struct {
	todo     *mydomain.Todo
	checkbox *widget.Bool
}

func NewTodoDelegate(todo *mydomain.Todo) *TodoDelegate {
	t := &TodoDelegate{
		todo:     todo,
		checkbox: new(widget.Bool),
	}
	t.checkbox.Value = todo.Completed
	return t
}

func (t *TodoDelegate) Event(ev event.Event) {
	if _, ok := ev.(system.FrameEvent); ok {
		if t.checkbox.Changed() {
			t.todo.Completed = t.checkbox.Value
			fmt.Printf("@ completed [%s] = %v\n", t.todo.Title, t.todo.Completed)
		}
	}
}

func (t TodoDelegate) String() string {
	return ">" + t.todo.Title
}

// --- Rendering -------------------------------------------------------------

func Todo(t *TodoDelegate) layout.Widget {
	title := t.todo.Title
	return func(gtx layout.Context) layout.Dimensions {
		//button := material.Button(Theme.Material(), clck.Clicker(), b.text)
		label := html.Div().Class("spaced").Class("level").Content(
			material.CheckBox(html.Theme.Material(), t.checkbox, status(t)).Layout,
			html.Glue(25, 2),
			html.P().Class("spaced").Text(title),
			html.Glue(3, 3),
		)
		return label(gtx)
	}
}

func status(t *TodoDelegate) string {
	if t.todo.Completed {
		return "done"
	}
	return "active"
}
