package fileutil

import (
	"io"
	"os"
)

// A file writer that opens the file on demand
// per default content is appended, later we might add other modes
type Writer struct {
	w    io.WriteCloser
	path string
	Mode os.FileMode // File mode for the file (set after open/create)
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

	if w.Mode != 0 {
		err = os.Chmod(w.path, w.Mode)
		if err != nil {
			return err
		}
	}

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