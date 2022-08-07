
logger
======

Simple levelled log library, using only builtin go libraries.

Usage
-----

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

