package html

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

func P() ParaStyler {
	return ParaStyler{}
}

type ParaStyler struct {
	color color.NRGBA
}

func (psty ParaStyler) Class(cssClass string) ParaStyler {
	psty.color = color.NRGBA{R: 230}
	return psty
}

func (psty ParaStyler) Text(txt string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		label := material.Body1(Theme, txt)
		if psty.color != noColor {
			label.Color = psty.color
		}
		return label.Layout(gtx)
	}
}
