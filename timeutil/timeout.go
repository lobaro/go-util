package timeutil

import "time"

// Keeps track of a Timeout, either a Task called "Done" before the timeout elapses
// or the onTimeout callback is called
type Timeout struct {
	timeout   time.Duration
	done      chan struct{}
	onTimeout func()
}

func NewTimeout(duration time.Duration, onTimeout func()) *Timeout {
	return &Timeout{
		timeout:   duration,
		done:      make(chan struct{}, 0),
		onTimeout: onTimeout,
	}
}

// Start the timeout, either Done() is called in time or onTimeout will be executed
func (to *Timeout) Start() {
	go func() {
		select {
		case <-to.done:
			return
		case <-time.After(to.timeout): // Erase on all pages takes 1.77s
			to.onTimeout()
		}
	}()
}

func (to *Timeout) Done() {
	close(to.done)
}