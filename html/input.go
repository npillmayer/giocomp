package html

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/npillmayer/giocomp/components"
	"github.com/npillmayer/giocomp/html/css"
)

// === Text Input ============================================================

func TextInput() _TextInput {
	return _TextInput{}
}

type _TextInput struct {
	css.Stylable
	text string
	hint string
}

func (ti _TextInput) Hint(hint string) _TextInput {
	ti.hint = hint
	return ti
}

func (ti _TextInput) Class(cssClass string) _TextInput {
	ti.Stylable = css.Apply(ti.Stylable, cssClass, Theme)
	return ti
}

func (ti _TextInput) Bind(editor *components.EditorDelegate) layout.Widget {
	myeditor := editor.Editor()
	hint := ti.hint
	return func(gtx layout.Context) layout.Dimensions {
		e := material.Editor(Theme.Material(), myeditor, ti.text)
		if hint != "" {
			e.Hint = hint
		}
		return ti.Styled(e.Layout)(gtx)
	}
}
