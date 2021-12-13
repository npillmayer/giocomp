package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/unit"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/examples/count/counter"
	"github.com/npillmayer/giocomp/html"
)

type myBusinessObject struct {
	count int
}

type myApplication struct {
	dom            *html.DOMLayout
	myBO           *myBusinessObject
	countUIControl *counter.CounterDelegate
}

func (myapp *myApplication) HandleEvent(ev event.Event) {
	myapp.dom.ForEvent(ev).Handle(myapp.countUIControl)
}

func (myapp *myApplication) Layout() {
	myapp.dom.Body().Content(
		html.Div().Class("spaced").Content(
			html.H1().Text("Counter"),
			html.P().Class("highlight").Text("This is an example app for a trivial counter"),
			html.Glue(30, 30),
			html.Div().Class("hbox").Content(
				counter.Counter(myapp.countUIControl),
				html.HFill(),
			),
		),
	)
}

// --- Main ------------------------------------------------------------------

func main() {
	myapp := myApplication{
		dom:  html.NewDOM(),
		myBO: &myBusinessObject{count: 1},
	}
	myapp.countUIControl = counter.New().Bind(&myapp.myBO.count)
	giocomp.Run(&myapp, app.Size(unit.Dp(350), unit.Dp(400)))
}
