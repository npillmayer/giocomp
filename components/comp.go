package components

import (
	"sync"
	"time"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/widget"
	"git.sr.ht/~gioverse/skel/scheduler"
)

type EventTarget interface {
	Event(event.Event)
}

type UpdateEvent struct {
	Value  interface{}
	Window *app.Window
}

// ImplementsEvent is a marker for type event.Event
func (uev UpdateEvent) ImplementsEvent() {}

// --- Clickable -------------------------------------------------------------

type Clickable struct {
	clicker widget.Clickable
}

func (clck *Clickable) Clicker() *widget.Clickable {
	return &clck.clicker
}

func (clck Clickable) Event(ev event.Event) {}

// --- Async -----------------------------------------------------------------

type InvalidateEvent bool

type Worker struct {
	mtx  sync.Mutex
	conn scheduler.Connection
}

func (w *Worker) Connect(conn scheduler.Connection) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	w.conn = conn
}

func (w *Worker) connection() (scheduler.Connection, bool) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	return w.conn, w.conn != nil
}

func (w *Worker) Async(f func()) {
	var conn scheduler.Connection
	var ok bool
	if conn, ok = w.connection(); !ok {
		return
	}
	conn.Schedule(func() interface{} {
		// Sleep to simulate expensive work like database
		// interactions, I/O, etc...
		time.Sleep(time.Millisecond * 500)
		f()
		return InvalidateEvent(true)
	})
}
