package view

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/view/css"
)

func P() _P {
	return _P{
		bodyText: material.Body1,
		color:    Theme.Material().Fg,
	}
}

type _P struct {
	css.Stylable
	bodyText func(*material.Theme, string) material.LabelStyle
	color    color.NRGBA
}

func (p _P) Class(cssClass string) _P {
	if cssClass == "em" {
		p.bodyText = material.Body2
	} else if cssClass == "highlight" {
		p.color = color.NRGBA{R: 180, A: 255}
	} else {
		p.Stylable = css.Apply(p.Stylable, cssClass, Theme)
	}
	return p
}

func (p _P) Text(txt string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		label := p.bodyText(Theme.Material(), txt)
		label.Color = p.color
		return p.Styled(layout.Widget(label.Layout))(gtx)
	}
}
