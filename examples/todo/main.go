package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/unit"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/examples/todo/mycomponents"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/view"
)

type myBusinessObject struct {
	todos *mydomain.TodoList
}

type myApplication struct {
	dom               *view.DOMLayout
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
		view.Div().Class("spaced").Content(
			view.H3().Text("Your ToDos"),
			view.P().Class("highlight").Text("This is an example app collecting ToDos"),
			view.Glue(30, 30),
			view.Div().Class("level").Content(
				mycomponents.TodoCreate(myapp.newToDoUIControl),
			),
			view.Glue(20, 20),
			view.Hr(),
			view.Glue(20, 20),
			mycomponents.Filter(myapp.filterUIControl),
			view.Glue(10, 10),
			mycomponents.TodoList(myapp.todoListUIControl),
		),
	)
}

// --- Main ------------------------------------------------------------------

func main() {
	myapp := myApplication{
		dom: view.NewDOM(),
		myBO: &myBusinessObject{
			todos: mydomain.NewTodoList(),
		},
	}
	myapp.todoListUIControl = mycomponents.NewTodoListDelegate(myapp.myBO.todos)
	myapp.newToDoUIControl = mycomponents.NewTodoCreator().Bind(myapp.todoListUIControl)
	myapp.filterUIControl = mycomponents.NewFilterDelegate().Bind(myapp.todoListUIControl)
	giocomp.Run(&myapp, app.Size(unit.Dp(400), unit.Dp(600)))
}
