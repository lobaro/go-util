package asyncio

import (
	"bytes"
	"github.com/Sirupsen/logrus"
	"io"
	"time"
)

type AsyncReader struct {
	r          io.Reader
	readCh     chan []byte   // Hands over data from read worker
	readBuffer *bytes.Buffer // Buffer data that is not requested by client

	readTimeout time.Duration
}

func NewAsyncReader(r io.Reader) io.Reader {
	ar := &AsyncReader{
		r:           r,
		readBuffer:  &bytes.Buffer{},
		readCh:      make(chan []byte, 0),
		readTimeout: 500 * time.Millisecond,
	}

	// Start the worker on create
	go ar.worker()
	return ar
}

func (ar *AsyncReader) SetReadTimeout(to time.Duration) {
	ar.readTimeout = to
}

func (ar *AsyncReader) ReadTimeout() time.Duration {
	return ar.readTimeout
}

func (ar *AsyncReader) Read(b []byte) (n int, err error) {
	// Data available, return it
	if ar.readBuffer.Len() > 0 {
		return ar.readFromBuffer(b)
	}

	// No data available, wait for new data with timeout
	select {
	case data := <-ar.readCh:
		ar.readBuffer.Write(data) // There might come more data as needed, so buffer it
		return ar.readFromBuffer(b)
	case <-time.After(ar.ReadTimeout()):
		return 0, io.EOF
	}
}

func (ar *AsyncReader) readFromBuffer(b []byte) (n int, err error) {
	n, err = ar.readBuffer.Read(b)
	return n, err
}

func (ar *AsyncReader) worker() {
	readBuf := make([]byte, 1024)
	for {
		n, err := ar.r.Read(readBuf) // Read from underlying reader
		if err != nil {
			// TODO: make the state available outside
			logrus.WithError(err).Error("AsyncReader worker died due to read error")
			return
		}

		data := make([]byte, n)
		copy(data, readBuf[:n])
		ar.readCh <- data
	}
}
