package broadcast

import (
	"sync"
	"fmt"
	"context"
)

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

func (b *Broadcaster) Close() {
	b.closed = true
	b.workers.Iter(func(w *worker) { w.Close() })
}

func (b *Broadcaster) Send(val interface{}) error {
	if b.closed {
		return fmt.Errorf("faild to send: broadcaster closed")
	}

	b.workers.Iter(func(w *worker) {
		if !w.closed {
			w.source <- val
		}
	})

	// TODO: remove closed workers!

	return nil
}


type HandlerFunc func(msg interface{})

// TODO: Return some handle to stop listening
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


func (s *threadSafeSlice) Iter(f func(*worker)) {
	s.Lock()
	defer s.Unlock()

	for _, worker := range s.workers {
		f(worker)
	}
}

type worker struct {
	source chan interface{}
	handler HandlerFunc
	quit   chan struct{}
	closed bool
}

func (w *worker) start() {
	w.source = make(chan interface{}, 10) // some buffer size to avoid blocking
	go func() {
		for {
			select {
			case msg := <-w.source:
				w.handler(msg)
			case <-w.quit:
				return
			}
		}
	}()
}



func (w *worker) Close() {
	w.closed = true
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

func (s *threadSafeSlice) Remove(w *worker) {
	s.Lock()
	defer s.Unlock()

	for i, worker := range s.workers {
		if w == worker {
			s.workers = append(s.workers[:i], s.workers[i+1:]...)
			return
		}
	}
}
