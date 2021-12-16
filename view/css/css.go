package css

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Box interface {
	ApplyDecorator(WidgetDecorator) Box
}

type CSS map[string]WidgetDecorator

func (css CSS) Apply(key string, box Box) Box {
	if styler, ok := css[key]; ok {
		box = box.ApplyDecorator(styler)
	}
	return box
}

type Theme struct {
	material *material.Theme
	styles   CSS
}

func (theme Theme) Material() *material.Theme {
	return theme.material
}

func (theme Theme) CSS() CSS {
	return theme.styles
}

func NewTheme(th *material.Theme, css CSS) Theme {
	if th == nil {
		th = material.NewTheme(gofont.Collection())
	}
	if css == nil {
		css = make(map[string]WidgetDecorator)
	}
	return Theme{
		material: th,
		styles:   css,
	}
}

// --- CSS Styling -----------------------------------------------------------

type Stylable struct {
	styler WidgetDecorator
}

// Apply widget style: wrap any existing styler into a new styler
func (sty Stylable) ApplyDecorator(ws WidgetDecorator) Box {
	if sty.styler == nil {
		sty.styler = ws
	} else {
		inner := sty.styler
		sty.styler = func(w layout.Widget) layout.Widget {
			return ws(inner(w))
		}
	}
	return sty
}

func (sty Stylable) Styled(w layout.Widget) layout.Widget {
	if sty.styler == nil {
		return w
	}
	return sty.styler(w)
}

type WidgetDecorator func(layout.Widget) layout.Widget

var _ Box = &Stylable{}

func Apply(sty Stylable, style string, theme Theme) Stylable {
	box := theme.styles.Apply(style, sty)
	if s, ok := box.(Stylable); ok {
		sty.styler = s.styler
	}
	return sty
}

// --- CSS style set ---------------------------------------------------------

type Styles struct {
	Fg          color.NRGBA
	Bg          color.NRGBA
	BorderColor color.NRGBA
	Border      int
	Rounded     bool
	Shaded      bool
}

func StylesFromTheme(theme Theme) Styles {
	return Styles{
		Fg:          theme.material.Fg,
		Bg:          theme.material.Bg,
		BorderColor: theme.material.Fg,
	}
}
