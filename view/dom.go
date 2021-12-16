package view

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/view/css"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// --- CSS Theming -----------------------------------------------------------

var materialTheme = material.NewTheme(gofont.Collection())

var Theme = css.NewTheme(materialTheme, css.CSS{
	"boxed": boxingStyler,
	"spaced": insetter(layout.Inset{
		Top:    unit.Dp(5),
		Right:  unit.Dp(10),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(10),
	}),
})

var noColor = color.NRGBA{}
var BlackColor = color.NRGBA{A: 255}

func boxingStyler(w layout.Widget) layout.Widget {
	border := widget.Border{
		Color: BlackColor,
		Width: unit.Dp(1),
	}
	return func(gtx layout.Context) layout.Dimensions {
		return border.Layout(gtx, w)
	}
}

func insetter(margins layout.Inset) css.WidgetDecorator {
	return func(w layout.Widget) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			return margins.Layout(gtx, w)
		}
	}
}

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

func (dom *DOMLayout) Body() *DOMLayout {
	return dom
}

func (dom *DOMLayout) Content(elements ...layout.Widget) {
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
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEnd,
	}.Layout(gtx, flexKids...)
	ev.Frame(&dom.ops)
}
