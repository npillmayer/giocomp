package mycomponents

import (
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

var defaultToDoLabelHint = "My new ToDo                                 "

type TodoCreatorDelegate struct {
	*components.Clickable
	editor *components.EditorDelegate
	components.Worker
	working  bool
	todo     *mydomain.Todo
	todolist *TodoListDelegate
}

func NewTodoCreator() *TodoCreatorDelegate {
	return &TodoCreatorDelegate{
		Clickable: &components.Clickable{},
		editor:    components.NewEditorDelegate(),
	}
}

func (c *TodoCreatorDelegate) Event(ev event.Event) {
	switch e := ev.(type) {
	case giocomp.ConnectEvent:
		c.Connect(e.Connection())
		return
	case system.FrameEvent:
		c.editor.Event(e)
	case components.UpdateEvent:
		// no update events expected for this component
	}
	if c.Clicker().Clicked() {
		if title := c.editor.Value(); title != "" {
			c.working = true
			c.Async(c.todolist, func() interface{} {
				c.todo = mydomain.NewTodo(title)
				c.editor.Editor().SetText("")
				c.working = false
				return c.todo
			})
		}
	}
}

func (c *TodoCreatorDelegate) Bind(list *TodoListDelegate) *TodoCreatorDelegate {
	c.todolist = list
	return c
}

func (c *TodoCreatorDelegate) Value() *mydomain.Todo {
	return c.todo
}

// --- Rendering -------------------------------------------------------------

func TodoCreate(c *TodoCreatorDelegate) layout.Widget {
	return html.Div().Class("level").Content(
		html.TextInput().Class("spaced").Class("boxed").Hint(defaultToDoLabelHint).Bind(c.editor),
		html.Button().Class("spaced").Class("is-primary").Text("Add").Bind(c.Clickable),
	)
}
