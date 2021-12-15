package components

import (
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/widget"
)

const voidText string = ""

type EditorDelegate struct {
	editor widget.Editor
	text   string
}

func NewEditorDelegate() *EditorDelegate {
	return &EditorDelegate{
		editor: widget.Editor{
			SingleLine: true,
			Submit:     true,
		},
	}
}

func (e *EditorDelegate) Bind(txt string) *EditorDelegate {
	e.text = txt // receive a domain object
	return e
}

func (e *EditorDelegate) Editor() *widget.Editor {
	return &e.editor
}

func (e EditorDelegate) Value() string {
	return e.text
}

func (e *EditorDelegate) Event(event event.Event) {
	switch event.(type) {
	case system.FrameEvent:
		for _, evt := range e.editor.Events() {
			if _, ok := evt.(widget.ChangeEvent); ok {
				e.text = e.editor.Text()
			}
			if ev, ok := evt.(widget.SubmitEvent); ok {
				e.text = ev.Text
			}
		}
	}
}
