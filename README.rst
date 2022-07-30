
logger
======

Simple levelled log library with no dependencies.
Go's builtin logger is goroutine-safe, and therefore so is this library.


Usage
-----

QuickStart
..........

A default logsetup logs to stderr. You can use the following.

.. code-block:: go

    import "willpittman.net/x/logger"

    logger.SetLevel(log.LvDebug)
    logger.Error("error msg")
    logger.Warn("warn msg")
    logger.Info("info msg")
    logger.Debug("debug msg")


You can also create your own loggers if you'd like.

.. code-block:: go

    import "willpittman.net/x/logger"

    # create logger
    logfile = os.Create("foo.log")
    defer logfile.Close()
    log = logger.New(logfile)

    # use logger
    log.SetLevel(log.LvDebug)
    log.Error("error msg")
    log.Warn("warn msg")
    log.Info("info msg")
    log.Debug("debug msg")


Testable Logs
.............

If you'd like to test logged messages, you may assign the `Logger` instance to a package variable.
In your test, you can then swap the logger with a `StubLogger` instance.
LogMessages are appended to a array fields on the struct - a separate array for each loglevel.

.. code-block:: go

    // internal/log/log.go

    package log
    import (
        "os"
        "willpittman.net/x/logger"
    )

    var Log logger.Interface
    func init() {
        logRaw := logger.New(os.Stderr)
        Log = &logRaw
    }

.. code-block:: go

    // main_test.go

    package main
    import "foo.com/yourpkg/internal/log/log"

    func TestMain(t *testing.T) {
        stubLog := logger.NewStubLogger()
        log.Log = &stubLog

        // code you are testing

        if len(stubLog.InfoMsgs) != 1 {
            t.Error("Expected a message to be logged")
        }
        if stubLog.InfoMsgs[0] != "Tadaa, I logged something" {
            t.Error("Expected the log message 'Tadaaa, I logged something'")
        }
    }

