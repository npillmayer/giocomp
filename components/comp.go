package components

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/widget"
)

type EventTarget interface {
	Event(event.Event)
}

type UpdateEvent struct {
	Value  interface{}
	Window *app.Window
}

// ImplementsEvent is a marker for type event.Event
func (uev UpdateEvent) ImplementsEvent() {}

// ---------------------------------------------------------------------------

type Clickable struct {
	clicker widget.Clickable
}

func (clck *Clickable) Clicker() *widget.Clickable {
	return &clck.clicker
}

func (clck Clickable) Event(ev event.Event) {}
