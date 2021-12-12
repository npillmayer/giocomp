package html

import (
	"strings"

	"gioui.org/layout"
)

func Div() _Div {
	div := _Div{}
	div.flex.Axis = layout.Vertical
	return div
}

type _Div struct {
	flex     layout.Flex
	elements []layout.FlexChild
}

func (div _Div) Class(cssClass string) _Div {
	if strings.Contains(cssClass, "hbox") {
		div.flex.Axis = layout.Horizontal
		div.flex.Alignment = layout.Middle
	}
	return div
}

func (div _Div) Content(elements ...layout.Widget) layout.Widget {
	div.elements = make([]layout.FlexChild, len(elements))
	for i, e := range elements {
		div.elements[i] = layout.Rigid(e)
	}
	return func(gtx layout.Context) layout.Dimensions {
		return div.flex.Layout(gtx, div.elements...)
	}
}
