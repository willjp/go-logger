package logger

import "regexp"

var leadingWhitespace = regexp.MustCompile(`(?m)(^\s+)`)
