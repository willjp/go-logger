package logger

import (
	"io"
	"os"
)

// Default Logger for the whole application.
var DefaultLogger Logger

// Set loglevel of DefaultLogger
func SetLevel(level LogLevel) {
	DefaultLogger.SetLevel(level)
}

// Set output of DefaultLogger
func SetOutput(w io.Writer) {
	DefaultLogger.SetOutput(w)
}

// Set format-flags of DefaultLogger
func SetFlags(flags int) {
	DefaultLogger.SetFlags(flags)
}

// Print debug message from DefaultLogger
func Debug(v ...interface{}) {
	DefaultLogger.callerDebug(v...)
}

// Print info message from DefaultLogger
func Info(v ...interface{}) {
	DefaultLogger.callerInfo(v...)
}

// Print warn message from DefaultLogger
func Warn(v ...interface{}) {
	DefaultLogger.callerWarn(v...)
}

// Print error message from DefaultLogger
func Error(v ...interface{}) {
	DefaultLogger.callerError(v...)
}

// Printf debug message from DefaultLogger
func Debugf(format string, v ...interface{}) {
	DefaultLogger.callerDebugf(format, v...)
}

// Printf info message from DefaultLogger
func Infof(format string, v ...interface{}) {
	DefaultLogger.callerInfof(format, v...)
}

// Printf warn message from DefaultLogger
func Warnf(format string, v ...interface{}) {
	DefaultLogger.callerWarnf(format, v...)
}

// Printf error message from DefaultLogger
func Errorf(format string, v ...interface{}) {
	DefaultLogger.callerErrorf(format, v...)
}

func init() {
	DefaultLogger = New(os.Stderr)
}
