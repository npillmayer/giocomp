package counter

import (
	"strconv"

	"gioui.org/layout"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/html"
)

type _Counter struct {
	components.Component
	count *int
}

func Counter() _Counter {
	return _Counter{}
}

func (c _Counter) Value() string {
	return strconv.Itoa(*c.count)
}

func (c _Counter) Bind(cnt *int) layout.Widget {
	c.count = cnt
	return html.P().Text(c.Value())
}
