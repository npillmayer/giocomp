package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/unit"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/examples/todo/mycomponents"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/html"
)

type myBusinessObject struct {
	todos *mydomain.TodoList
}

type myApplication struct {
	dom               *html.DOMLayout
	myBO              *myBusinessObject
	newToDoUIControl  *mycomponents.TodoCreatorDelegate
	todoListUIControl *mycomponents.TodoListDelegate
}

func (myapp *myApplication) HandleEvent(ev event.Event) {
	myapp.dom.ForEvent(ev).Handle(
		myapp.newToDoUIControl,
		myapp.todoListUIControl,
	)
}

func (myapp *myApplication) Layout() {
	myapp.dom.Body().Content(
		html.Div().Class("spaced").Content(
			html.H2().Text("Your To-Dos"),
			html.P().Class("highlight").Text("This is an example app collecting To Dos"),
			html.Glue(30, 30),
			html.Div().Class("level").Content(
				mycomponents.TodoCreate(myapp.newToDoUIControl),
			),
			mycomponents.TodoList(myapp.todoListUIControl),
		),
	)
}

// --- Main ------------------------------------------------------------------

func main() {
	myapp := myApplication{
		dom: html.NewDOM(),
		myBO: &myBusinessObject{
			todos: mydomain.NewTodoList(),
		},
	}
	myapp.todoListUIControl = mycomponents.NewTodoListDelegate(myapp.myBO.todos)
	myapp.newToDoUIControl = mycomponents.NewTodoCreator().Bind(myapp.todoListUIControl)
	giocomp.Run(&myapp, app.Size(unit.Dp(400), unit.Dp(600)))
}
