package view

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/view/css"
)

type HeadingStyler struct {
	css.Stylable
	level int
	color color.NRGBA
}

func H1() HeadingStyler {
	return HeadingStyler{level: 1}
}

func H2() HeadingStyler {
	return HeadingStyler{level: 2}
}

func H3() HeadingStyler {
	return HeadingStyler{level: 3}
}

func Title() HeadingStyler {
	return HeadingStyler{level: 0}
}

func (h HeadingStyler) Class(cssClass string) HeadingStyler {
	if cssClass == "highlight" {
		h.color = color.NRGBA{R: 180, A: 255}
	} else {
		h.Stylable = css.Apply(h.Stylable, cssClass, Theme)
	}
	return h
}

func (h HeadingStyler) Text(txt string) layout.Widget {
	level := h.level
	color := h.color
	heading := func(gtx layout.Context) layout.Dimensions {
		var label material.LabelStyle
		switch level {
		case 0:
			label = material.H1(Theme.Material(), txt)
		case 1:
			label = material.H2(Theme.Material(), txt)
		case 2:
			label = material.H3(Theme.Material(), txt)
		case 3:
			label = material.H4(Theme.Material(), txt)
		default:
			label = material.H5(Theme.Material(), txt)
		}
		if color != noColor {
			label.Color = color
		}
		return label.Layout(gtx)
	}
	return h.Styled(heading)
}
