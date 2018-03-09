package async

import (
	"time"
	"context"
)

type Worker struct {
	ctx    context.Context
	cancel context.CancelFunc
	doWork DoWorkFn
	err    error
	done   chan struct{} // Closed when done
	Delay  time.Duration // Delay between doWork calls
}

type DoWorkFn func() error

func NewWorker(fn DoWorkFn) *Worker {
	return &Worker{
		doWork: fn,
		Delay:  1 * time.Second,
	}
}

// Cancel stops the worker
func (w *Worker) Cancel() {
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
}

// Done is closed when the work loop ended
func (w *Worker) Done() <-chan struct{} {
	return w.done
}

// Start starts the worker, does nothing when already running
func (w *Worker) Start() {
	if w.cancel != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	w.ctx = ctx
	w.cancel = cancel
	w.err = nil
	w.done = make(chan struct{}, 0)

	go w.doWorkLoop(ctx)
}

func (w *Worker) Error() error {
	return w.err
}

func (w *Worker) doWorkLoop(ctx context.Context) {
	defer func() {
		close(w.done)
	}()
	for {
		err := w.doWork()
		if err != nil {
			w.err = err
			w.Cancel()
			return
		}

		select {
		case <-w.ctx.Done():
			return
		case <-time.After(w.Delay):
			// Do more work
		}

	}
}
