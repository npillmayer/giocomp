package mycomponents

import (
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/html"
)

// --- Behaviour -------------------------------------------------------------

type FilterDelegate struct {
	suppressCompleted bool
	checkbox          *widget.Bool
	list              *TodoListDelegate
}

func NewFilterDelegate() *FilterDelegate {
	return &FilterDelegate{
		checkbox: new(widget.Bool),
	}
}

func (f *FilterDelegate) Bind(list *TodoListDelegate) *FilterDelegate {
	f.list = list
	return f
}

func (f *FilterDelegate) Event(ev event.Event) {
	if _, ok := ev.(system.FrameEvent); ok {
		if f.checkbox.Changed() {
			f.suppressCompleted = f.checkbox.Value
			if f.list != nil {
				f.list.FilterCompleted(f.suppressCompleted)
			}
		}
	}
}

// --- Rendering -------------------------------------------------------------

func Filter(t *FilterDelegate) layout.Widget {
	filter := material.Switch(html.Theme.Material(), t.checkbox, "")
	return html.Div().Class("level").Content(
		filter.Layout,
		html.P().Class("spaced").Text("hide completed"),
	)
}
