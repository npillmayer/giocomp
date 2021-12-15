package html

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/npillmayer/giocomp/html/css"
)

// Adapted copy from:
// https://git.sr.ht/~eliasnaur/gio-example/tree/main/item/outlay/fan/widget/boring/

// --- Bordered panel with rounded corners -----------------------------------

type PanelStyle struct {
	css.Styles
	cornerRadius float32
	Alignment    layout.Alignment
	Content      layout.Widget
}

func (panelStyle PanelStyle) WithStyles(styles css.Styles) PanelStyle {
	panelStyle.Styles = styles
	if styles.Rounded {
		panelStyle.cornerRadius = 8
	} else {
		panelStyle.cornerRadius = 0
	}
	return panelStyle
}

func (panelStyle PanelStyle) Wrap(w layout.Widget) layout.Widget {
	panelStyle.Content = w
	return func(gtx layout.Context) layout.Dimensions {
		return panelStyle.Layout(gtx, panelStyle.Content)
	}
}

func (p *PanelStyle) Layout(gtx C, w layout.Widget) D {
	var outerRadius, innerRadius float32
	border := float32(p.Border)
	if p.Rounded {
		outerRadius = min(p.cornerRadius*2, float32(gtx.Constraints.Max.X))
		outerRadius = min(outerRadius, float32(gtx.Constraints.Max.Y))
		outerRadius *= 0.5
		if outerRadius < 4*border {
			outerRadius = 0
		} else {
			innerRadius = outerRadius - border
		}
	}
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return Rect{
				Color:  p.BorderColor,
				Size:   layout.FPt(gtx.Constraints.Max),
				Radius: outerRadius,
			}.Layout(gtx)
		}),
		layout.Stacked(func(gtx C) D {
			return layout.UniformInset(unit.Dp(border)).Layout(gtx, func(gtx C) D {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx C) D {
						return Rect{
							Color:  p.Bg,
							Size:   layout.FPt(gtx.Constraints.Max),
							Radius: innerRadius,
						}.Layout(gtx)
					}),
					layout.Stacked(func(gtx C) D {
						return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx C) D {
							gtx.Constraints.Min = gtx.Constraints.Max
							layout.Center.Layout(gtx, func(gtx C) D {
								dims := w(gtx)
								if p.Alignment != layout.Middle {
									dims.Size.X = gtx.Constraints.Max.X
								}
								return dims
							})
							return D{Size: gtx.Constraints.Max}
						})
					}),
				)
			})
		}),
	)
}

// --- Rect with rounded corners ---------------------------------------------

// Rect creates a rectangle of the provided background color with
// Dimensions specified by size and a corner radius (on all corners)
// specified by radii.
type Rect struct {
	Color  color.NRGBA
	Size   f32.Point
	Radius float32
}

// Layout renders the Rect into the provided context
func (r Rect) Layout(gtx C) D {
	return DrawRect(gtx, r.Color, r.Size, r.Radius)
}

// DrawRect creates a rectangle of the provided background color with
// Dimensions specified by size and a corner radius (on all corners)
// specified by radii.
func DrawRect(gtx C, background color.NRGBA, size f32.Point, radius float32) D {
	bounds := f32.Rectangle{Max: size}
	paint.FillShape(gtx.Ops, background, clip.UniformRRect(bounds, radius).Op(gtx.Ops))
	return layout.Dimensions{Size: image.Pt(int(size.X), int(size.Y))}
}

// ---------------------------------------------------------------------------

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
