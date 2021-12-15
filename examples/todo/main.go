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
	filterUIControl   *mycomponents.FilterDelegate
}

func (myapp *myApplication) HandleEvent(ev event.Event) {
	myapp.dom.ForEvent(ev).Handle(
		myapp.newToDoUIControl,
		myapp.todoListUIControl,
		myapp.filterUIControl,
	)
}

func (myapp *myApplication) Layout() {
	myapp.dom.Body().Content(
		html.Div().Class("spaced").Content(
			html.H3().Text("Your ToDos"),
			html.P().Class("highlight").Text("This is an example app collecting ToDos"),
			html.Glue(30, 30),
			html.Div().Class("level").Content(
				mycomponents.TodoCreate(myapp.newToDoUIControl),
			),
			html.Glue(20, 20),
			html.Hr(),
			html.Glue(20, 20),
			mycomponents.Filter(myapp.filterUIControl),
			html.Glue(10, 10),
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
	myapp.filterUIControl = mycomponents.NewFilterDelegate().Bind(myapp.todoListUIControl)
	giocomp.Run(&myapp, app.Size(unit.Dp(400), unit.Dp(600)))
}
