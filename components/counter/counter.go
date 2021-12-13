package counter

import (
	"fmt"
	"strconv"

	"gioui.org/io/event"
	"gioui.org/layout"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

type CounterDelegate struct {
	*components.Clickable
	components.Worker
	count *int
}

func New() *CounterDelegate {
	var count int // allocate a default domain object
	return &CounterDelegate{
		count:     &count,
		Clickable: &components.Clickable{},
	}
}

func (c *CounterDelegate) Event(ev event.Event) {
	if c.Clicker().Clicked() {
		fmt.Printf("@ counter received event %#v\n", ev)
		*c.count++
	}
	switch e := ev.(type) {
	case giocomp.ConnectEvent:
		c.Connect(e.Connection())
	}
}

func (c *CounterDelegate) Bind(cnt *int) *CounterDelegate {
	c.count = cnt // receive a domain object
	return c
}

func (c *CounterDelegate) Value() string {
	return strconv.Itoa(*c.count) // do something display-related to domain object
}

// --- Rendering -------------------------------------------------------------

func Counter(c *CounterDelegate) layout.Widget {
	return html.Div().Class("hbox").Class("boxed").Content(
		html.H2().Class("spaced").Text(c.Value()),
		html.Button().Class("spaced").Text("Count").Bind(c.Clickable),
	)
}
