package logutil

import (
	"github.com/sirupsen/logrus"
	"io"
	"slices"
	"sync"
)

// LogWriteHook implements the logrus.Hook interface so it can be added for easy file logging (via ThreadsafeWriter) via logrus.AddHook.
// It needs to be created with NewWriterHook first.
type LogWriteHook struct {
	target    ThreadsafeWriter
	formatter logrus.Formatter
	levels    []logrus.Level
}

var defaultFormatter = &logrus.TextFormatter{
	ForceColors:   false,
	DisableColors: true,
}

// NewWriterHook creates a new LogWriteHook to the selected io.Writer. format and levels can be nil if the default TextFormatter with no colours and all log levels should be used
func NewWriterHook(target io.Writer, format logrus.Formatter, levels []logrus.Level) LogWriteHook {
	tWriter := ThreadsafeWriter{
		Writer: target,
		Mutex:  &sync.Mutex{},
	}
	lHook := LogWriteHook{
		target: tWriter,
	}

	if format != nil {
		lHook.formatter = format
	} else {
		lHook.formatter = defaultFormatter
	}

	// output all levels if levels == nil
	if levels != nil {
		lHook.levels = levels
	} else {
		lHook.levels = logrus.AllLevels
	}

	return lHook
}

func (lf *LogWriteHook) Levels() []logrus.Level {
	return lf.levels
}

func (lf *LogWriteHook) Fire(entry *logrus.Entry) error {
	message, err := lf.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = lf.target.Write(message)

	return err
}

func (lf *LogWriteHook) SetFormatter(format logrus.Formatter) {
	lf.formatter = format
}

func (lf *LogWriteHook) SetLevels(levels []logrus.Level) {
	lf.levels = levels
}

func (lf *LogWriteHook) AddLevels(levels []logrus.Level) {
	lf.levels = append(lf.levels, levels...)
}

func (lf *LogWriteHook) RemoveLevels(levels []logrus.Level) {
	newLvls := make([]logrus.Level, 0)
	for _, lvl := range lf.levels {
		if !slices.Contains(levels, lvl) {
			newLvls = append(newLvls, lvl)
		}
	}
	lf.levels = newLvls
}
