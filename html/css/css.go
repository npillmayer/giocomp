package css

import (
	"fmt"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Box interface {
	ApplyStyler(WidgetStyler) Box
}

type CSS map[string]WidgetStyler

func (css CSS) Apply(key string, box Box) Box {
	if styler, ok := css[key]; ok {
		fmt.Printf("@ applying style with key %q\n", key)
		box = box.ApplyStyler(styler)
	} else {
		fmt.Printf("@ style for key %q is void\n", key)
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
		css = make(map[string]WidgetStyler)
	}
	return Theme{
		material: th,
		styles:   css,
	}
}

// --- CSS Styling -----------------------------------------------------------

type Stylable struct {
	styler WidgetStyler
}

// Apply widget style: wrap any existing styler into a new styler
func (sty Stylable) ApplyStyler(ws WidgetStyler) Box {
	fmt.Printf("@ in Styleable.ApplyStyler, styler = %#v\n", sty.styler)
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

type WidgetStyler func(layout.Widget) layout.Widget

var _ Box = &Stylable{}

func Apply(sty Stylable, style string, theme Theme) Stylable {
	box := theme.styles.Apply(style, sty)
	if s, ok := box.(Stylable); ok {
		sty.styler = s.styler
	}
	return sty
}