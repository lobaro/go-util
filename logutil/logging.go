package logutil

import (
	"io"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/fatih/color"
)

var Output = &ThreadsafeWriter{
	Writer: color.Output,
	Mutex:  &sync.Mutex{},
}

type ThreadsafeWriter struct {
	Writer io.Writer
	Mutex  *sync.Mutex
}

func (w ThreadsafeWriter) Write(p []byte) (n int, err error) {
	w.Mutex.Lock()
	n, err = w.Writer.Write(p)
	w.Mutex.Unlock()
	return
}

// Setup default logging.
// Logrus with color output
func SetupLogging() {
	logrus.SetOutput(Output)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})

	logrus.Info("Setup default logging")
}
