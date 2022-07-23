package logger

import (
	"reflect"
	"testing"
)

func TestStubLoggerImplementsInterface(t *testing.T) {
	belongsToInterface := func(Interface) bool {
		return true
	}
	logger_ := NewStubLogger()
	if !belongsToInterface(&logger_) {
		t.Error("Logger does not conform to Interface")
	}
}

func TestLogMessagesRecorded(t *testing.T) {
	t.Run("Debug", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvDebug)
		msg := "foo"
		logger_.Debug(msg)
		if len(logger_.DebugMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.DebugMsgs[0] != msg {
			t.Errorf("Expected '%s', Received '%s'", msg, logger_.DebugMsgs[0])
		}
	})

	t.Run("Info", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvInfo)
		msg := "foo"
		logger_.Info(msg)
		if len(logger_.InfoMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.InfoMsgs[0] != msg {
			t.Errorf("Expected '%s', Received '%s'", msg, logger_.InfoMsgs[0])
		}
	})

	t.Run("Warn", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvWarn)
		msg := "foo"
		logger_.Warn(msg)
		if len(logger_.WarnMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.WarnMsgs[0] != msg {
			t.Errorf("Expected '%s', Received '%s'", msg, logger_.WarnMsgs[0])
		}
	})

	t.Run("Error", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvError)
		msg := "foo"
		logger_.Error(msg)
		if len(logger_.ErrorMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.ErrorMsgs[0] != msg {
			t.Errorf("Expected '%s', Received '%s'", msg, logger_.ErrorMsgs[0])
		}
	})

	t.Run("Debugf", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvDebug)
		val := "foo"
		logger_.Debugf("val: %s", val)
		expects := "val: foo"
		if len(logger_.DebugMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.DebugMsgs[0] != expects {
			t.Errorf("Expected '%s', Received '%s'", expects, logger_.DebugMsgs[0])
		}
	})

	t.Run("Infof", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvInfo)
		val := "foo"
		logger_.Infof("val: %s", val)
		expects := "val: foo"
		if len(logger_.InfoMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.InfoMsgs[0] != expects {
			t.Errorf("Expected '%s', Received '%s'", expects, logger_.InfoMsgs[0])
		}
	})

	t.Run("Warnf", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvWarn)
		val := "foo"
		logger_.Warnf("val: %s", val)
		expects := "val: foo"
		if len(logger_.WarnMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.WarnMsgs[0] != expects {
			t.Errorf("Expected '%s', Received '%s'", expects, logger_.WarnMsgs[0])
		}
	})

	t.Run("Errorf", func(t *testing.T) {
		logger_ := NewStubLogger()
		logger_.SetLevel(LvError)
		val := "foo"
		logger_.Errorf("val: %s", val)
		expects := "val: foo"
		if len(logger_.ErrorMsgs) != 1 {
			t.Error("No Log Message Received")
		}
		if logger_.ErrorMsgs[0] != expects {
			t.Errorf("Expected '%s', Received '%s'", expects, logger_.ErrorMsgs[0])
		}
	})
}

func TestStubLoggerLevels(t *testing.T) {
	t.Run("SetLevel works with Regular Logs", func(t *testing.T) {
		tcases := []struct {
			test  string
			level LogLevel
			logs  []string
		}{
			{
				test:  "LvNone output",
				level: LvNone,
				logs:  []string{},
			},
			{
				test:  "LvError output",
				level: LvError,
				logs:  []string{"error"},
			},
			{
				test:  "LvWarn output",
				level: LvWarn,
				logs:  []string{"error", "warn"},
			},
			{
				test:  "LvInfo output",
				level: LvInfo,
				logs:  []string{"error", "warn", "info"},
			},
			{
				test:  "LvDebug output",
				level: LvDebug,
				logs:  []string{"error", "warn", "info", "debug"},
			},
		}

		for _, tcase := range tcases {
			t.Run(tcase.test, func(t *testing.T) {
				logger_ := NewStubLogger()

				logger_.SetLevel(tcase.level)
				logger_.Error("error")
				logger_.Warn("warn")
				logger_.Info("info")
				logger_.Debug("debug")

				logs := []string{}
				logs = append(logs, logger_.ErrorMsgs...)
				logs = append(logs, logger_.WarnMsgs...)
				logs = append(logs, logger_.InfoMsgs...)
				logs = append(logs, logger_.DebugMsgs...)
				if !reflect.DeepEqual(tcase.logs, logs) {
					t.Errorf("Log Messages do not match.\nExpected:\n'%s'\nReceived:\n'%s'", tcase.logs, logs)
				}
			})
		}

	})

	t.Run("SetLevel works with Format Logs", func(t *testing.T) {
		tcases := []struct {
			test  string
			level LogLevel
			logs  []string
		}{
			{
				test:  "LvNone output",
				level: LvNone,
				logs:  []string{},
			},
			{
				test:  "LvError output",
				level: LvError,
				logs:  []string{"error"},
			},
			{
				test:  "LvWarn output",
				level: LvWarn,
				logs:  []string{"error", "warn"},
			},
			{
				test:  "LvInfo output",
				level: LvInfo,
				logs:  []string{"error", "warn", "info"},
			},
			{
				test:  "LvDebug output",
				level: LvDebug,
				logs:  []string{"error", "warn", "info", "debug"},
			},
		}

		for _, tcase := range tcases {
			t.Run(tcase.test, func(t *testing.T) {
				logger_ := NewStubLogger()

				logger_.SetLevel(tcase.level)
				logger_.Errorf("%s", "error")
				logger_.Warnf("%s", "warn")
				logger_.Infof("%s", "info")
				logger_.Debugf("%s", "debug")

				logs := []string{}
				logs = append(logs, logger_.ErrorMsgs...)
				logs = append(logs, logger_.WarnMsgs...)
				logs = append(logs, logger_.InfoMsgs...)
				logs = append(logs, logger_.DebugMsgs...)
				if !reflect.DeepEqual(tcase.logs, logs) {
					t.Errorf("Log Messages do not match.\nExpected:\n'%s'\nReceived:\n'%s'", tcase.logs, logs)
				}
			})
		}

	})
}
