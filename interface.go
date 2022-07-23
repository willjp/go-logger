package logger

import "io"

type Interface interface {
	SetLevel(level LogLevel)
	SetOutput(w io.Writer)
	SetFlags(flags int)
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}
