package broadcast

import (
	"context"
	"fmt"
	"sync"
)

type HandlerFunc func(msg interface{})

type Broadcaster struct {
	workers    threadSafeSlice
	globalQuit chan struct{}
	closed     bool
}

func NewBoradcaster() *Broadcaster {
	return &Broadcaster{
		workers: threadSafeSlice{
			workers: make([]*worker, 0),
		},
		globalQuit: make(chan struct{}),
	}
}

func (b *Broadcaster) StartListen(handler HandlerFunc) context.CancelFunc {
	w := &worker{}
	w.handler = handler
	w.start()
	b.workers.Append(w)

	return func() {
		w.Close()
		b.workers.Remove(w)
	}
}

func (b *Broadcaster) Send(val interface{}) error {
	if b.closed {
		return fmt.Errorf("failed to send: broadcaster closed")
	}

	closedWorkers := make([]*worker, 0)

	b.workers.Iter(func(w *worker) {
		if !w.closed {
			w.source <- val
		} else {
			closedWorkers = append(closedWorkers, w)
		}
	})

	// Remove closed workers
	for _, w := range closedWorkers {
		b.workers.Remove(w)
	}

	return nil
}

func (b *Broadcaster) Close() {
	if !b.closed {
		b.closed = true
		b.workers.Iter(func(w *worker) { w.Close() })
		close(b.globalQuit) // Needed?
	}
}

type worker struct {
	source  chan interface{}
	handler HandlerFunc
	quit    chan struct{} // Used from extern, needed?
	closeCh chan struct{} // Used internal on close
	closed  bool
}

func (w *worker) start() {
	w.source = make(chan interface{}, 10) // some buffer size to avoid blocking
	w.closeCh = make(chan struct{}, 0)
	go func() {
		for {
			select {
			case msg := <-w.source:
				w.handler(msg)
			case <-w.quit:
				return
			case <-w.closeCh:
				return
			}
		}
	}()
}

func (w *worker) Close() {
	if !w.closed {
		w.closed = true
		close(w.closeCh)
	}
}

type threadSafeSlice struct {
	sync.Mutex
	workers []*worker
}

func (s *threadSafeSlice) Append(w *worker) {
	s.Lock()
	defer s.Unlock()

	s.workers = append(s.workers, w)
}

func (s *threadSafeSlice) Remove(w *worker) int {
	s.Lock()
	defer s.Unlock()

	for i, worker := range s.workers {
		if w == worker {
			s.workers = append(s.workers[:i], s.workers[i+1:]...)
			return i
		}
	}
	return -1
}

func (s *threadSafeSlice) Iter(f func(*worker)) {
	s.Lock()
	defer s.Unlock()

	for _, worker := range s.workers {
		f(worker)
	}
}
