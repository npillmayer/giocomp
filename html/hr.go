package html

import (
	"gioui.org/f32"
	"gioui.org/layout"
)

func Hr() layout.Widget {
	return func(gtx C) D {
		return Rect{
			Color: Theme.Material().Fg,
			Size:  f32.Pt(float32(gtx.Constraints.Max.X), 2),
		}.Layout(gtx)
	}
}
