// Simple reusable logger implementation that supports loglevels.
//
// For convenience, DefaultLogger is created and logs to STDERR.
// Package Functions modify/use this DefaultLogger.
//
//   Ex.
//       logger.SetLevel(logger.LvDebug)
//       logger.Debug("Did something risky")
//
//
// You can also create/customize your own loggers using the same interface
//   Ex.
//       logfile = os.Create("foo.log")
//       defer logfile.Close()
//       mylogger := logger.New(logfile)
//       mylogger.SetLevel(logger.LvDebug)
//       mylogger.Debug("Did something risky")
//
// Inspired by xorm's logger.
//
//
package logger
