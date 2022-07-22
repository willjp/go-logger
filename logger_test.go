package logger

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

var leadingWhitespace = regexp.MustCompile(`(?m)(^\s+)`)

func TestNew(t *testing.T) {
	t.Run("Default Log Prefix", func(t *testing.T) {
		var prefix string
		writer := strings.Builder{}
		logger_ := New(&writer)

		prefix = "[ERROR] "
		if logger_.error.Prefix() != prefix {
			t.Errorf("Unexpected logger.error.Prefix(). expected: '%s', received: '%s'",
				prefix, logger_.error.Prefix())
		}

		prefix = "[INFO ] "
		if logger_.info.Prefix() != prefix {
			t.Errorf("Unexpected logger.info.Prefix(). expected: '%s', received: '%s'",
				prefix, logger_.info.Prefix())
		}

		prefix = "[WARN ] "
		if logger_.warn.Prefix() != prefix {
			t.Errorf("Unexpected logger.warn.Prefix(). expected: '%s', received: '%s'",
				prefix, logger_.warn.Prefix())
		}

		prefix = "[DEBUG] "
		if logger_.debug.Prefix() != prefix {
			t.Errorf("Unexpected logger.debug.Prefix(). expected: '%s', received: '%s'",
				prefix, logger_.debug.Prefix())
		}
	})

	t.Run("Default Log Flags", func(t *testing.T) {
		format := log.Ldate | log.Ltime | log.Llongfile
		writer := strings.Builder{}
		logger_ := New(&writer)

		if logger_.error.Flags() != format {
			t.Errorf("Unexpected logger.error.Flags(). expected: '%d', received: '%d'",
				format, logger_.error.Flags())
		}

		if logger_.info.Flags() != format {
			t.Errorf("Unexpected logger.info.Flags(). expected: '%d', received: '%d'",
				format, logger_.info.Flags())
		}

		if logger_.warn.Flags() != format {
			t.Errorf("Unexpected logger.warn.Flags(). expected: '%d', received: '%d'",
				format, logger_.warn.Flags())
		}

		if logger_.debug.Flags() != format {
			t.Errorf("Unexpected logger.debug.Flags(). expected: '%d', received: '%d'",
				format, logger_.debug.Flags())
		}
	})
}

func TestSetOutput(t *testing.T) {
	oldWriter := strings.Builder{}
	newWriter := strings.Builder{}
	logger_ := New(&oldWriter)
	logger_.SetLevel(LvDebug)
	logger_.SetFlags(0)

	logger_.SetOutput(&newWriter)
	expects := leadingWhitespace.ReplaceAllString(
		`[ERROR] error
		 [WARN ] warn
		 [INFO ] info
		 [DEBUG] debug
		`,
		"",
	)
	logger_.Error("error")
	logger_.Warn("warn")
	logger_.Info("info")
	logger_.Debug("debug")
	if oldWriter.String() != "" {
		t.Errorf("oldWriter was written to")
	}
	if newWriter.String() != expects {
		t.Errorf("newWriter did not receive expected logs\nexpected:\n'%s'\nreceived:'%s'", expects, newWriter.String())
	}

}

func TestLogger(t *testing.T) {
	t.Run("SetLevel works with Regular Logs", func(t *testing.T) {
		tcases := []struct {
			test  string
			level LogLevel
			logs  string
		}{
			{
				test:  "LvNone output",
				level: LvNone,
				logs:  "",
			},
			{
				test:  "LvError output",
				level: LvError,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error
					`,
					"",
				),
			},
			{
				test:  "LvWarn output",
				level: LvWarn,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error
					 [WARN ] warn
					`,
					"",
				),
			},
			{
				test:  "LvInfo output",
				level: LvInfo,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error
					 [WARN ] warn
					 [INFO ] info
					`,
					"",
				),
			},
			{
				test:  "LvDebug output",
				level: LvDebug,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error
					 [WARN ] warn
					 [INFO ] info
					 [DEBUG] debug
					`,
					"",
				),
			},
		}

		for _, tcase := range tcases {
			t.Run(tcase.test, func(t *testing.T) {
				writer := strings.Builder{}
				logger_ := New(&writer)
				logger_.SetFlags(0)

				logger_.SetLevel(tcase.level)
				logger_.Error("error")
				logger_.Warn("warn")
				logger_.Info("info")
				logger_.Debug("debug")
				if writer.String() != tcase.logs {
					t.Errorf("Log Messages do not match.\nExpected:\n'%s'\nReceived:\n'%s'", tcase.logs, writer.String())
				}
			})
		}

	})

	t.Run("SetLevel works with Format Logs", func(t *testing.T) {
		tcases := []struct {
			test  string
			level LogLevel
			logs  string
		}{
			{
				test:  "LvNone output",
				level: LvNone,
				logs:  "",
			},
			{
				test:  "LvError output",
				level: LvError,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error: foo
					`,
					"",
				),
			},
			{
				test:  "LvWarn output",
				level: LvWarn,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error: foo
					[WARN ] warn: foo
					`,
					"",
				),
			},
			{
				test:  "LvInfo output",
				level: LvInfo,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error: foo
					[WARN ] warn: foo
					[INFO ] info: foo
					`,
					"",
				),
			},
			{
				test:  "LvDebug output",
				level: LvDebug,
				logs: leadingWhitespace.ReplaceAllString(
					`[ERROR] error: foo
					[WARN ] warn: foo
					[INFO ] info: foo
					[DEBUG] debug: foo
					`,
					"",
				),
			},
		}

		for _, tcase := range tcases {
			t.Run(tcase.test, func(t *testing.T) {
				writer := strings.Builder{}
				logger_ := New(&writer)
				logger_.SetLevel(tcase.level)
				logger_.SetFlags(0)
				logger_.Errorf("error: %s", "foo")
				logger_.Warnf("warn: %s", "foo")
				logger_.Infof("info: %s", "foo")
				logger_.Debugf("debug: %s", "foo")
				if writer.String() != tcase.logs {
					t.Errorf("Log Messages do not match.\nExpected:\n'%s'\nReceived:\n'%s'", tcase.logs, writer.String())
				}
			})
		}
	})
}
