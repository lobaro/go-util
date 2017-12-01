package broadcast

import (
	"testing"
	"bytes"
	"time"
	"sync"
)

// Takes bytes from source and publish it via bc
func publish(t *testing.T, source *bytes.Buffer, bc *Broadcaster) {
	buf := make([]byte, 1)
	for {
		n, err := source.Read(buf)
		if err != nil || n == 0 {
			return
		}
		msg := make([]byte, n)
		copy(msg, buf[:n])
		t.Logf("Publish %v", msg)
		bc.Send(msg)
	}

	t.Logf("Publier done.")
}

func TestBroadcastCancel(t *testing.T) {
	source := &bytes.Buffer{}
	bc := NewBoradcaster()

	received := &bytes.Buffer{}

	go publish(t, source, bc)

	wg := sync.WaitGroup{}
	wg.Add(1)
	cancel := bc.StartListen(func(msg interface{}) {
		t.Logf("Received: %v", msg)
		received.Write(msg.([]byte))
		wg.Done()
	})

	source.WriteByte(1)
	wg.Wait() // Wait to be published
	cancel()
	source.WriteByte(2)
	time.Sleep(100 * time.Millisecond) // Give some time to receive

	if received.Len() != 1 {
		t.Errorf("Expected received 1 bytes but got %d", received.Len())
	}

	if len(bc.workers.workers) != 0 {
		t.Errorf("Expected 0 workers but got %d", len(bc.workers.workers))
	}

}


func TestBroadcastMultiple(t *testing.T) {
	source := &bytes.Buffer{}
	bc := NewBoradcaster()

	received := &bytes.Buffer{}

	go publish(t, source, bc)

	wg := sync.WaitGroup{}
	wg.Add(2)
	bc.StartListen(func(msg interface{}) {
		t.Logf("Received (1): %v", msg)
		received.Write(msg.([]byte))
		wg.Done()
	})

	bc.StartListen(func(msg interface{}) {
		t.Logf("Received (2): %v", msg)
		received.Write(msg.([]byte))
		wg.Done()
	})

	source.WriteByte(1)
	wg.Wait() // Wait to be published

	if received.Len() != 2 {
		t.Errorf("Expected received 2 bytes but got %d", received.Len())
	}

	if len(bc.workers.workers) != 2 {
		t.Errorf("Expected 2 workers but got %d", len(bc.workers.workers))
	}

}
