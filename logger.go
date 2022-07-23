package logger

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	level LogLevel
	flags int
	error *log.Logger
	info  *log.Logger
	warn  *log.Logger
	debug *log.Logger
}

// Create a new custom Logger
func New(writer io.Writer) Logger {
	return Logger{
		level: defaultLogLevel,
		error: log.New(writer, "[ERROR] ", defaultLogFlags),
		info:  log.New(writer, "[INFO ] ", defaultLogFlags),
		warn:  log.New(writer, "[WARN ] ", defaultLogFlags),
		debug: log.New(writer, "[DEBUG] ", defaultLogFlags),
	}
}

func (l *Logger) Flags() int {
	return l.flags
}

func (l *Logger) Level() LogLevel {
	return l.level
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) SetOutput(w io.Writer) {
	l.error.SetOutput(w)
	l.info.SetOutput(w)
	l.warn.SetOutput(w)
	l.debug.SetOutput(w)
}

func (l *Logger) SetFlags(flags int) {
	l.flags = flags
	l.error.SetFlags(flags)
	l.info.SetFlags(flags)
	l.warn.SetFlags(flags)
	l.debug.SetFlags(flags)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level >= LvDebug {
		l.debug.Output(2, fmt.Sprint(v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level >= LvInfo {
		l.info.Output(2, fmt.Sprint(v...))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level >= LvWarn {
		l.warn.Output(2, fmt.Sprint(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level >= LvError {
		l.error.Output(2, fmt.Sprint(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level >= LvDebug {
		l.debug.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level >= LvInfo {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level >= LvWarn {
		l.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level >= LvError {
		l.error.Output(2, fmt.Sprintf(format, v...))
	}
}

// Following methods omit caller's call-stack when logging
func (l *Logger) callerDebug(v ...interface{}) {
	if l.level >= LvDebug {
		l.debug.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) callerInfo(v ...interface{}) {
	if l.level >= LvInfo {
		l.info.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) callerWarn(v ...interface{}) {
	if l.level >= LvWarn {
		l.warn.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) callerError(v ...interface{}) {
	if l.level >= LvError {
		l.error.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) callerDebugf(format string, v ...interface{}) {
	if l.level >= LvDebug {
		l.debug.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) callerInfof(format string, v ...interface{}) {
	if l.level >= LvInfo {
		l.info.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) callerWarnf(format string, v ...interface{}) {
	if l.level >= LvWarn {
		l.warn.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) callerErrorf(format string, v ...interface{}) {
	if l.level >= LvError {
		l.error.Output(3, fmt.Sprintf(format, v...))
	}
}
