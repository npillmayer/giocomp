package components

import (
	"sync"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget"
	"git.sr.ht/~gioverse/skel/scheduler"
)

// --- Clickable -------------------------------------------------------------

type Clickable struct {
	clicker widget.Clickable
}

func (clck *Clickable) Clicker() *widget.Clickable {
	return &clck.clicker
}

func (clck Clickable) Event(ev event.Event) {}

// --- Async -----------------------------------------------------------------

type EventTarget interface {
	Event(event.Event)
}

type UpdateEvent struct {
	Value       interface{}
	Target      EventTarget
	Invalidates bool
}

// ImplementsEvent is a marker for type event.Event
func (uev UpdateEvent) ImplementsEvent() {}

type InvalidateEvent bool

// ImplementsEvent is a marker for type event.Event
func (ev InvalidateEvent) ImplementsEvent() {}

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

func (w *Worker) Async(target EventTarget, f func() interface{}) {
	var conn scheduler.Connection
	var ok bool
	if conn, ok = w.connection(); !ok {
		return
	}
	conn.Schedule(func() interface{} {
		return UpdateEvent{
			Value:       f(),
			Target:      target,
			Invalidates: true,
		}
	})
}

// --- Disabled state --------------------------------------------------------

func Enable(cond bool, w layout.Widget) layout.Widget {
	if w == nil {
		return nil
	}
	if cond {
		return w
	}
	inner := w
	return func(gtx layout.Context) layout.Dimensions {
		//gtx.Queue = nil
		gtx = gtx.Disabled()
		return inner(gtx)
	}
}
