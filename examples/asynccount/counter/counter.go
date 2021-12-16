package counter

import (
	"strconv"
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/view"
)

// --- Behaviour -------------------------------------------------------------

type CounterDelegate struct {
	*components.Clickable
	components.Worker
	working bool
	count   *int
}

func New() *CounterDelegate {
	var count int // allocate a default domain object
	return &CounterDelegate{
		count:     &count,
		Clickable: &components.Clickable{},
	}
}

func (c *CounterDelegate) Event(ev event.Event) {
	switch e := ev.(type) {
	case giocomp.ConnectEvent:
		c.Connect(e.Connection())
		return
	case components.UpdateEvent:
		if e.Target == c {
			*c.count = e.Value.(int)
			c.working = false
		}
	}
	if c.Clicker().Clicked() {
		c.working = true
		c.Async(c, func() interface{} {
			// Sleep to simulate expensive work like database
			// interactions, I/O, etc...
			time.Sleep(time.Millisecond * 1200)
			return *c.count + 1
		})
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
	counter := view.Div().Class("level").Class("boxed").Content(
		view.H2().Class("spaced").Text(c.Value()),
		view.Button().Class("spaced").Class("is-primary").Text("Count").Bind(c.Clickable),
	)
	return components.Enable(!c.working, counter)
}
