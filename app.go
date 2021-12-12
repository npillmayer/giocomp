package giocomp

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"git.sr.ht/~gioverse/skel/scheduler"
	"git.sr.ht/~gioverse/skel/window"
	"github.com/npillmayer/giocomp/components/counter"
	"github.com/npillmayer/giocomp/html"
)

func Run() {
	bus := scheduler.NewWorkerPool()
	w := window.NewWindower(bus)
	go func() {
		// Wait for all windows to close, then exit.
		w.Run()
		os.Exit(0)
	}()
	//window.NewWindow(bus, loop)
	window.NewWindow(bus, MainLoop)
	app.Main()
}

func MainLoop(w *app.Window, conn scheduler.Connection) error {
	dom := html.NewDOM(op.Ops{}, material.NewTheme(gofont.Collection()))
	var mycount int = 77
	for {
		select {
		case event := <-w.Events():
			switch event := event.(type) {
			case system.DestroyEvent:
				return event.Err
			case system.FrameEvent:
				// Lay out the UI here
				dom.ForEvent(event).Body(
					html.H1().Text("Hello"),
					html.H2().Text("World"),
					counter.Counter().Bind(&mycount),
				)
				//,
				//counter.Counter().Bind(&mycount))
			}
		case update := <-conn.Output():
			// Handle any requests to modify the window that came over the bus.
			window.Update(w, update)
			// Check for application state updates and handle them.
			switch update := update.(type) {
			case NewNumberEvent:
				_ = int(update) // TODO
				w.Invalidate()
			}
		}
	}
}

type NewNumberEvent int

func loop(w *app.Window, conn scheduler.Connection) error {
	th := material.NewTheme(gofont.Collection())
	var (
		ops            op.Ops
		add, newWindow widget.Clickable
	)
	currentNumber := 5

	for {
		select {
		case event := <-w.Events():
			switch event := event.(type) {
			case system.DestroyEvent:
				return event.Err
			case system.FrameEvent:
				if add.Clicked() {
					// Schedule asynchronous work to fetch an new number. In real applications,
					// these are often expensive and blocking operations.
					conn.Schedule(func() interface{} {
						// Sleep to simulate expensive work like database
						// interactions, I/O, etc...
						time.Sleep(time.Millisecond * 200)
						return NewNumberEvent(rand.Intn(100) + 1)
					})
				}
				if newWindow.Clicked() {
					// Launch a new window running another copy of this event
					// loop.
					window.NewWindow(conn, loop)
				}
				// Lay out the UI here
				gtx := layout.NewContext(&ops, event)
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(
						material.H1(th, strconv.Itoa(currentNumber)).Layout,
					),
					layout.Rigid(material.Button(th, &add, "Add").Layout),
					layout.Rigid(material.Button(th, &newWindow, "New Window").Layout),
				)
				event.Frame(&ops)
			}
		case update := <-conn.Output():
			// Handle any requests to modify the window that came over the bus.
			window.Update(w, update)
			// Check for application state updates and handle them.
			switch update := update.(type) {
			case NewNumberEvent:
				currentNumber += int(update)
				w.Invalidate()
			}
		}
	}
}
