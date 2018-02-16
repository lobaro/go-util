package asyncio

import (
	"bytes"
	"io"
	"time"
)

type ReadErrorHandler func (err error, errCnt int) bool

type AsyncReader struct {
	r          io.Reader     // underlying reader
	readCh     chan []byte   // Hands over data from read worker
	readBuffer *bytes.Buffer // Buffer data that is not read by client

	readTimeout time.Duration
	error       error
	readErrorHandler ReadErrorHandler
}

func NewAsyncReader(r io.Reader) *AsyncReader {
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

func (ar *AsyncReader) SetReadErrorHandler(cb ReadErrorHandler) {
	ar.readErrorHandler = cb
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
		return ar.readBuffer.Read(b)
	}

	// No data available, wait for new data with timeout
	select {
	case data := <-ar.readCh:
		ar.readBuffer.Write(data) // There might come more data as needed, so buffer it
		return ar.readBuffer.Read(b)
	case <-time.After(ar.ReadTimeout()):
		return 0, io.EOF
	}
}

func (ar *AsyncReader) Error() error {
	return ar.error
}

func (ar *AsyncReader) worker() {
	errCnt := 0
	readBuf := make([]byte, 1024)
	for {
		n, err := ar.r.Read(readBuf) // Read from underlying reader
		if err != nil {
			ar.error = err
			if !ar.readErrorHandler(err, errCnt) {
				return
			}

			errCnt++
			continue
		}

		errCnt = 0
		data := make([]byte, n)
		copy(data, readBuf[:n])
		ar.readCh <- data
	}
}
