package logger

import (
	"fmt"
	"io"
)

// StubLogger records the strings that were logged to each log-level.
// It is designed for testing log messages.
type StubLogger struct {
	level     LogLevel
	flags     int
	ErrorMsgs []string
	InfoMsgs  []string
	WarnMsgs  []string
	DebugMsgs []string
}

// Creates a StubLogger
// Since this will most likely be used to test for log messages, it is created with the highest possible loglevel (all messages will be written).
func NewStubLogger() StubLogger {
	return StubLogger{
		level:     LvDebug,
		flags:     defaultLogFlags,
		ErrorMsgs: []string{},
		InfoMsgs:  []string{},
		WarnMsgs:  []string{},
		DebugMsgs: []string{},
	}
}

func (l *StubLogger) Flags() int {
	return l.flags
}

func (l *StubLogger) Level() LogLevel {
	return l.level
}

func (l *StubLogger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *StubLogger) SetOutput(w io.Writer) {
	// ignored, the purpose of this object is to record calls
	return
}

func (l *StubLogger) SetFlags(flags int) {
	l.flags = flags
}

func (l *StubLogger) Debug(v ...interface{}) {
	if l.level >= LvDebug {
		l.DebugMsgs = append(l.DebugMsgs, fmt.Sprint(v...))
	}
}

func (l *StubLogger) Info(v ...interface{}) {
	if l.level >= LvInfo {
		l.InfoMsgs = append(l.InfoMsgs, fmt.Sprint(v...))
	}
}

func (l *StubLogger) Warn(v ...interface{}) {
	if l.level >= LvWarn {
		l.WarnMsgs = append(l.WarnMsgs, fmt.Sprint(v...))
	}
}

func (l *StubLogger) Error(v ...interface{}) {
	if l.level >= LvError {
		l.ErrorMsgs = append(l.ErrorMsgs, fmt.Sprint(v...))
	}
}

func (l *StubLogger) Debugf(format string, v ...interface{}) {
	if l.level >= LvDebug {
		l.DebugMsgs = append(l.DebugMsgs, fmt.Sprintf(format, v...))
	}
}

func (l *StubLogger) Infof(format string, v ...interface{}) {
	if l.level >= LvInfo {
		l.InfoMsgs = append(l.InfoMsgs, fmt.Sprintf(format, v...))
	}
}

func (l *StubLogger) Warnf(format string, v ...interface{}) {
	if l.level >= LvWarn {
		l.WarnMsgs = append(l.WarnMsgs, fmt.Sprintf(format, v...))
	}
}

func (l *StubLogger) Errorf(format string, v ...interface{}) {
	if l.level >= LvError {
		l.ErrorMsgs = append(l.ErrorMsgs, fmt.Sprintf(format, v...))
	}
}
