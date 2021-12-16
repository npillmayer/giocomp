package view

import (
	"image"

	"gioui.org/layout"
)

func Glue(w, h int) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		max := image.Point{X: w, Y: h}
		max = gtx.Constraints.Constrain(max)
		return layout.Dimensions{
			Size: max,
		}
	}
}

func HFill() layout.Widget {
	return Glue(32000, 1)
}

func VFill() layout.Widget {
	return Glue(1, 32000)
}
