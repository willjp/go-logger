package logger

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestDefaultLoggerSetOutput(t *testing.T) {
	newWriter := strings.Builder{}
	SetLevel(LvDebug)
	SetFlags(0)

	SetOutput(&newWriter)
	expects := leadingWhitespace.ReplaceAllString(
		`[ERROR] error
		 [WARN ] warn
		 [INFO ] info
		 [DEBUG] debug
		`,
		"",
	)
	Error("error")
	Warn("warn")
	Info("info")
	Debug("debug")
	if newWriter.String() != expects {
		t.Errorf("newWriter did not receive expected logs\nexpected:\n'%s'\nreceived:'%s'", expects, newWriter.String())
	}
}

func TestDefaultLogger(t *testing.T) {
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
				SetOutput(&writer)
				SetFlags(0)

				SetLevel(tcase.level)
				Error("error")
				Warn("warn")
				Info("info")
				Debug("debug")
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
				SetOutput(&writer)
				SetLevel(tcase.level)
				SetFlags(0)
				Errorf("error: %s", "foo")
				Warnf("warn: %s", "foo")
				Infof("info: %s", "foo")
				Debugf("debug: %s", "foo")
				if writer.String() != tcase.logs {
					t.Errorf("Log Messages do not match.\nExpected:\n'%s'\nReceived:\n'%s'", tcase.logs, writer.String())
				}
			})
		}
	})

	t.Run("References file that logged message", func(t *testing.T) {
		writer := strings.Builder{}
		SetOutput(&writer)
		SetLevel(LvDebug)
		SetFlags(log.Lshortfile)
		Error("error")
		Warn("warn")
		Info("info")
		Debug("debug")
		output := writer.String()

		prefixRx := regexp.MustCompile(`(?m)(^\s*\[[a-zA-Z ]+\] )`)
		output = prefixRx.ReplaceAllString(output, "")

		shortfileRx := regexp.MustCompile(`^default_test.go:[0-9]+: `)
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			if !shortfileRx.MatchString(line) {
				t.Errorf("Does not use log-caller's callstack. Received: %s", line)
			}
		}
	})
}
