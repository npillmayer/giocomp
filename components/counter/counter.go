package counter

import (
	"fmt"
	"strconv"

	"gioui.org/io/event"
	"gioui.org/layout"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

type CounterDelegate struct {
	*components.Clickable
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
}

func (c *CounterDelegate) Bind(cnt *int) *CounterDelegate {
	c.count = cnt // receive a domain object
	return c
}

func (c CounterDelegate) Value() string {
	return strconv.Itoa(*c.count) // do something display-related to domain object
}

// --- Rendering -------------------------------------------------------------

func Counter(c *CounterDelegate) layout.Widget {
	return html.Div().Class("hbox").Content(
		html.H2().Text(c.Value()),
		html.Button().Class("is-primary").Text("Count").Bind(c.Clickable),
	)
}
