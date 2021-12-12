package html

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type HeadingStyler struct {
	level int
	color color.NRGBA
}

func H1() HeadingStyler {
	return HeadingStyler{level: 1}
}

func H2() HeadingStyler {
	return HeadingStyler{level: 2}
}

func Title() HeadingStyler {
	return HeadingStyler{level: 0}
}

func (lsty HeadingStyler) Class(cssClass string) HeadingStyler {
	lsty.color = color.NRGBA{R: 230}
	return lsty
}

func (lsty HeadingStyler) Text(txt string) layout.Widget {
	level := lsty.level
	color := lsty.color
	return func(gtx layout.Context) layout.Dimensions {
		var label material.LabelStyle
		switch level {
		case 0:
			label = material.H1(Theme, txt)
		case 1:
			label = material.H2(Theme, txt)
		case 2:
			label = material.H3(Theme, txt)
		default:
			label = material.H4(Theme, txt)
		}
		if color != noColor {
			label.Color = color
		}
		return label.Layout(gtx)
	}
}
