package html

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/components"
)

var Theme = material.NewTheme(gofont.Collection())

var noColor = color.NRGBA{}

/*
type Canvas struct {
	theme *material.Theme
	gtx   layout.Context
}

func (canvas Canvas) Theme() *material.Theme {
	return canvas.theme
}

func (canvas Canvas) Context() layout.Context {
	return canvas.gtx
}
*/

//type Element func(canvas Canvas) layout.Dimensions

// --- DOM -------------------------------------------------------------------

type DOMLayout struct {
	ops   op.Ops
	event event.Event
}

func NewDOM() *DOMLayout {
	return &DOMLayout{
		ops: op.Ops{},
	}
}

func (dom *DOMLayout) ForEvent(ev event.Event) *DOMLayout {
	dom.event = ev
	return dom
}

func (dom *DOMLayout) Handle(targets ...components.EventTarget) *DOMLayout {
	for _, t := range targets {
		t.Event(dom.event)
	}
	return dom
}

func (dom *DOMLayout) Body(elements ...layout.Widget) {
	ev, ok := dom.event.(system.FrameEvent)
	if !ok {
		return
	}
	gtx := layout.NewContext(&dom.ops, ev)
	flexKids := make([]layout.FlexChild, len(elements))
	for i, e := range elements {
		flexKids[i] = layout.Rigid(e)
	}
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, flexKids...)
	ev.Frame(&dom.ops)
}
