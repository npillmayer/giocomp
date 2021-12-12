package main

import (
	"gioui.org/io/event"
	"github.com/npillmayer/giocomp"
	"github.com/npillmayer/giocomp/components/counter"
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
	myapp.dom.Body(
		html.H1().Text("Counter"),
		html.P().Text("This is an example app for a trivial counter"),
		counter.Counter(myapp.countUIControl),
	)
}

// --- Main ------------------------------------------------------------------

func main() {
	myapp := myApplication{
		dom:  html.NewDOM(),
		myBO: &myBusinessObject{count: 1},
	}
	myapp.countUIControl = counter.New().Bind(&myapp.myBO.count)
	giocomp.Run(&myapp)
}
