package fileutil

import (
	"io"
)

// A file writer that opens the file on demand
// per default content is appended, later we might add other modes
type Writer struct {
	w    io.WriteCloser
	path string
}

func (w *Writer) open() error {
	if w.w != nil {
		w.w.Close()
	}

	logFile, err := OpenCreateAppend(w.path)
	if err != nil {
		w.w = nil
		return err
	}
	w.w = logFile

	return nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.w == nil {
		err := w.open()
		if err != nil {
			return 0, err
		}
	}
	return w.w.Write(p)
}

// when closing Writer it can be re-opend by just calling Write
func (w *Writer) Close() error {
	if w.w != nil {
		err := w.w.Close()
		w.w = nil
		return err
	}
	return nil
}

func NewWriter(filename string) (*Writer) {
	return &Writer{
		path: filename,
	}
}