package giocomp

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"git.sr.ht/~gioverse/skel/scheduler"
	"git.sr.ht/~gioverse/skel/window"
	"github.com/npillmayer/giocomp/components"
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
	runner.application.HandleEvent(ConnectEvent{conn})
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
			switch u := update.(type) {
			case components.InvalidateEvent:
				if u {
					w.Invalidate()
				}
			case components.UpdateEvent:
				runner.application.HandleEvent(u) // handle component events
				if u.Invalidates {
					w.Invalidate()
				}
			}
		}
	}
}

func Run(application Application, opts ...app.Option) {
	bus := scheduler.NewWorkerPool()
	w := window.NewWindower(bus)
	go func() {
		// Wait for all windows to close, then exit.
		w.Run()
		os.Exit(0)
	}()
	runner := ApplicationRunner{application: application}
	window.NewWindow(bus, runner.MainLoop, opts...)
	app.Main()
}

type ConnectEvent struct {
	conn scheduler.Connection
}

// ImplementsEvent is a marker interface in io.event.
func (cev ConnectEvent) ImplementsEvent() {}

func (cev ConnectEvent) Connection() scheduler.Connection {
	return cev.conn
}
