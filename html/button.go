package html

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/html/css"
)

func Button() _Button {
	return _Button{}
}

type _Button struct {
	css.Stylable
	isPrimary bool
	text      string
}

func (b _Button) Class(cssClass string) _Button {
	if cssClass == "is-primary" {
		b.isPrimary = true
	} else {
		fmt.Printf("@ applying CSS class %q on BUTTON\n", cssClass)
		b.Stylable = css.Apply(b.Stylable, cssClass, Theme)
	}
	return b
}

func (b _Button) Text(txt string) _Button {
	b.text = txt
	return b
}

func (b _Button) Bind(clck *components.Clickable) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		button := material.Button(Theme.Material(), clck.Clicker(), b.text)
		if !b.isPrimary {
			button.Background = color.NRGBA{R: 130, G: 130, B: 130, A: 0xff}
		}
		return b.Styled(button.Layout)(gtx)
	}
}
