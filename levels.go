package logger

type LogLevel int8

// Enum of LogLevels.
const (
	LvNone LogLevel = 10 * iota
	LvError
	LvWarn
	LvInfo
	LvDebug
)
