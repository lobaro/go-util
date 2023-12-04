package logutil

import (
	"io"
	"sync"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
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

	logrus.Debug("Setup default logging")
}

func SetupJsonLogging() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Debug("Setup json logging")
}

func NewDefaultLogger() *logrus.Logger {
	l := logrus.New()
	l.Out = Output
	l.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	}

	return l
}
