package mycomponents

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/examples/todo/mydomain"
	"github.com/npillmayer/giocomp/html"
	"github.com/npillmayer/giocomp/html/css"
)

// --- Behaviour -------------------------------------------------------------

type TodoDelegate struct {
	todo     *mydomain.Todo
	checkbox *widget.Bool
}

func NewTodoDelegate(todo *mydomain.Todo) *TodoDelegate {
	t := &TodoDelegate{
		todo:     todo,
		checkbox: new(widget.Bool),
	}
	t.checkbox.Value = todo.Completed
	return t
}

func (t *TodoDelegate) Event(ev event.Event) {
	if _, ok := ev.(system.FrameEvent); ok {
		if t.checkbox.Changed() {
			t.todo.Completed = t.checkbox.Value
		}
	}
}

// --- Rendering -------------------------------------------------------------

func Todo(t *TodoDelegate) layout.Widget {
	title := t.todo.Title
	return func(gtx layout.Context) layout.Dimensions {
		label := html.Div().Class("spaced").Class("level").Content(
			opaque(material.CheckBox(html.Theme.Material(), t.checkbox, "")),
			html.Glue(12, 2),
			html.P().Text(title),
		)
		return panel(label, t)(gtx)
	}
}

var (
	colorInactive       = color.NRGBA{R: 0xbb, G: 0xbb, B: 0xbb, A: 255}
	colorBorderInactive = color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 255}
	colorActive         = color.NRGBA{R: 0x99, G: 0xaa, B: 0xba, A: 255}
	colorBorderActive   = color.NRGBA{R: 0x55, G: 0x88, B: 0x98, A: 255}
)

func panel(w layout.Widget, t *TodoDelegate) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		cssStyle := css.StylesFromTheme(html.Theme)
		cssStyle.Rounded = true
		panelStyle := html.PanelStyle{}.WithStyles(cssStyle)
		gtx.Constraints.Max.Y = 120
		panelStyle.Alignment = layout.Start
		panelStyle.Border = 1
		if t.todo.Completed {
			panelStyle.Bg = colorInactive
			panelStyle.BorderColor = colorBorderInactive
		} else {
			panelStyle.Bg = colorActive
			panelStyle.BorderColor = colorBorderActive
		}
		return layout.Inset{
			Top:    unit.Dp(2),
			Bottom: unit.Dp(2),
		}.Layout(gtx, panelStyle.Wrap(w))
	}
}

// --- Don't look here -------------------------------------------------------

// This is an awful hack, but I'm not in the mood of diving into the checkbox/semantic widget
// stuff right now
func opaque(cb material.CheckBoxStyle) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		stack := layout.Stack{Alignment: layout.Center}
		return stack.Layout(gtx,
			layout.Stacked(
				func(gtx layout.Context) layout.Dimensions {
					dims := html.Rect{
						Color: html.Theme.Material().Bg,
						Size:  f32.Pt(33, 33),
					}.Layout(gtx)
					dims.Size.X += 7
					return dims
				},
			),
			layout.Stacked(cb.Layout),
		)
	}
}
