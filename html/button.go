package html

import (
	"image/color"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/components"
)

func Button() ButtonStyler {
	return ButtonStyler{}
}

type ButtonStyler struct {
	isPrimary bool
	text      string
}

func (bsty ButtonStyler) Class(cssClass string) ButtonStyler {
	bsty.isPrimary = strings.Contains(cssClass, "is-primary")
	return bsty
}

func (bsty ButtonStyler) Text(txt string) ButtonStyler {
	bsty.text = txt
	return bsty
}

func (bsty ButtonStyler) Bind(clck *components.Clickable) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		button := material.Button(Theme, clck.Clicker(), bsty.text)
		if !bsty.isPrimary {
			button.Background = color.NRGBA{R: 120, G: 120, B: 120, A: 0xff}
		}
		return button.Layout(gtx)
	}
}
