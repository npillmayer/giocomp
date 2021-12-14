package mycomponents

import (
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
	return &TodoDelegate{
		todo:     todo,
		checkbox: new(widget.Bool),
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
	if t.todo.IsDone() {
		return "done"
	}
	return "active"
}
