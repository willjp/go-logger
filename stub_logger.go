package logger

import (
	"fmt"
	"io"

	"willpittman.net/x/logger/internal/spinlock"
)

// StubLogger records the strings that were logged to each log-level.
// It is designed for testing log messages.
//
// StubLogger is threadsafe.
type StubLogger struct {
	level LogLevel
	flags int

	ErrorMsgs []string
	InfoMsgs  []string
	WarnMsgs  []string
	DebugMsgs []string

	optsLock  *spinlock.SpinLock
	errorLock *spinlock.SpinLock
	infoLock  *spinlock.SpinLock
	warnLock  *spinlock.SpinLock
	debugLock *spinlock.SpinLock
}

// Creates a StubLogger
// Since this will most likely be used to test for log messages, it is created with the highest possible loglevel (all messages will be written).
func NewStubLogger() StubLogger {
	return StubLogger{
		level: LvDebug,
		flags: defaultLogFlags,

		ErrorMsgs: []string{},
		InfoMsgs:  []string{},
		WarnMsgs:  []string{},
		DebugMsgs: []string{},

		optsLock:  &spinlock.SpinLock{},
		errorLock: &spinlock.SpinLock{},
		infoLock:  &spinlock.SpinLock{},
		warnLock:  &spinlock.SpinLock{},
		debugLock: &spinlock.SpinLock{},
	}
}

func (this *StubLogger) Flags() int {
	this.optsLock.Acquire()
	defer this.optsLock.Release()
	return this.flags
}

func (this *StubLogger) Level() LogLevel {
	this.optsLock.Acquire()
	defer this.optsLock.Release()
	return this.level
}

func (this *StubLogger) SetLevel(level LogLevel) {
	this.optsLock.Acquire()
	defer this.optsLock.Release()
	this.level = level
}

func (this *StubLogger) SetOutput(w io.Writer) {
	// ignored, the purpose of this object is to record calls
	return
}

func (this *StubLogger) SetFlags(flags int) {
	this.optsLock.Acquire()
	defer this.optsLock.Release()
	this.flags = flags
}

func (this *StubLogger) Debug(v ...interface{}) {
	if this.Level() >= LvDebug {
		this.debugLock.Acquire()
		defer this.debugLock.Release()
		this.DebugMsgs = append(this.DebugMsgs, fmt.Sprint(v...))
	}
}

func (this *StubLogger) Info(v ...interface{}) {
	if this.Level() >= LvInfo {
		this.infoLock.Acquire()
		defer this.infoLock.Release()
		this.InfoMsgs = append(this.InfoMsgs, fmt.Sprint(v...))
	}
}

func (this *StubLogger) Warn(v ...interface{}) {
	if this.Level() >= LvWarn {
		this.warnLock.Acquire()
		defer this.warnLock.Release()
		this.WarnMsgs = append(this.WarnMsgs, fmt.Sprint(v...))
	}
}

func (this *StubLogger) Error(v ...interface{}) {
	if this.Level() >= LvError {
		this.errorLock.Acquire()
		defer this.errorLock.Release()
		this.ErrorMsgs = append(this.ErrorMsgs, fmt.Sprint(v...))
	}
}

func (this *StubLogger) Debugf(format string, v ...interface{}) {
	if this.Level() >= LvDebug {
		this.debugLock.Acquire()
		defer this.debugLock.Release()
		this.DebugMsgs = append(this.DebugMsgs, fmt.Sprintf(format, v...))
	}
}

func (this *StubLogger) Infof(format string, v ...interface{}) {
	if this.Level() >= LvInfo {
		this.infoLock.Acquire()
		defer this.infoLock.Release()
		this.InfoMsgs = append(this.InfoMsgs, fmt.Sprintf(format, v...))
	}
}

func (this *StubLogger) Warnf(format string, v ...interface{}) {
	if this.Level() >= LvWarn {
		this.warnLock.Acquire()
		defer this.warnLock.Release()
		this.WarnMsgs = append(this.WarnMsgs, fmt.Sprintf(format, v...))
	}
}

func (this *StubLogger) Errorf(format string, v ...interface{}) {
	if this.Level() >= LvError {
		this.errorLock.Acquire()
		defer this.errorLock.Release()
		this.ErrorMsgs = append(this.ErrorMsgs, fmt.Sprintf(format, v...))
	}
}
