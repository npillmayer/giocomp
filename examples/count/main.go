package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/unit"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/examples/count/counter"
	"github.com/npillmayer/giocomp/view"
)

type myBusinessObject struct {
	count int
}

type myApplication struct {
	dom            *view.DOMLayout
	myBO           *myBusinessObject
	countUIControl *counter.CounterDelegate
}

func (myapp *myApplication) HandleEvent(ev event.Event) {
	myapp.dom.ForEvent(ev).Handle(myapp.countUIControl)
}

func (myapp *myApplication) Layout() {
	myapp.dom.Body().Content(
		view.Div().Class("spaced").Content(
			view.H1().Text("Counter"),
			view.P().Class("highlight").Text("This is an example app for a trivial counter"),
			view.Glue(30, 30),
			view.Div().Class("level").Content(
				counter.Counter(myapp.countUIControl),
				view.HFill(),
			),
		),
	)
}

// --- Main ------------------------------------------------------------------

func main() {
	myapp := myApplication{
		dom:  view.NewDOM(),
		myBO: &myBusinessObject{count: 1},
	}
	myapp.countUIControl = counter.New().Bind(&myapp.myBO.count)
	giocomp.Run(&myapp, app.Size(unit.Dp(350), unit.Dp(400)))
}
