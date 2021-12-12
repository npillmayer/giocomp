package html

import (
	"fmt"
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
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
	event system.FrameEvent
	//theme *material.Theme
}

func NewDOM(ops op.Ops, theme *material.Theme) *DOMLayout {
	fmt.Println("@@@ new DOM")
	return &DOMLayout{
		ops: ops,
		//theme: theme,
	}
}

func (dom *DOMLayout) ForEvent(event system.FrameEvent) *DOMLayout {
	dom.event = event
	fmt.Println("@@@ for event")
	return dom
}

func (dom *DOMLayout) Body(elements ...layout.Widget) {
	gtx := layout.NewContext(&dom.ops, dom.event)
	s := make([]layout.FlexChild, len(elements))
	for i, e := range elements {
		s[i] = layout.Rigid(e)
	}
	fmt.Println("@@@ BODY")
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		// layout.Rigid(
		// 	material.H1(Theme, "Hello").Layout,
		// ),
		s...,
	)
	dom.event.Frame(&dom.ops)
}
