package giocomp

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"git.sr.ht/~gioverse/skel/scheduler"
	"git.sr.ht/~gioverse/skel/window"
)

// Example modified from:
// https://git.sr.ht/~gioverse/skel/tree/main/item/example/readme/windower/windower.go

type Application interface {
	HandleEvent(event.Event)
	Layout()
}

type ApplicationRunner struct {
	application Application
}

func (runner ApplicationRunner) MainLoop(w *app.Window, conn scheduler.Connection) error {
	for {
		select {
		case event := <-w.Events():
			switch event := event.(type) {
			case system.DestroyEvent:
				return event.Err
			case system.FrameEvent:
				runner.application.HandleEvent(event) // handle frame events
				runner.application.Layout()           // then lay out UI tree
			}
		case update := <-conn.Output():
			// handle any requests to modify the window that came over the bus.
			window.Update(w, update)
		}
	}
}

func Run(application Application) {
	bus := scheduler.NewWorkerPool()
	w := window.NewWindower(bus)
	go func() {
		// Wait for all windows to close, then exit.
		w.Run()
		os.Exit(0)
	}()
	runner := ApplicationRunner{application: application}
	window.NewWindow(bus, runner.MainLoop)
	app.Main()
}

/*
func MainLoop(w *app.Window, conn scheduler.Connection) error {
	dom := html.NewDOM()
	var countDomainObject int = 1
	countDelegate := counter.New().Bind(&countDomainObject)
	for {
		select {
		case event := <-w.Events():
			switch event := event.(type) {
			case system.DestroyEvent:
				return event.Err
			case system.FrameEvent:
				// handle events
				dom.ForEvent(event).Handle(countDelegate)
				// lay out UI tree
				dom.Body(
					html.H1().Text("Counter"),
					html.P().Text("This is an example app for a trivial counter"),
					counter.Counter(countDelegate),
				)
			}
		case update := <-conn.Output():
			// handle any requests to modify the window that came over the bus.
			window.Update(w, update)
		}
	}
}
*/
