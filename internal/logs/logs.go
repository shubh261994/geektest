package logs

import (
	"runtime/debug"
)

func Error(args ...interface{}) {
	debug.PrintStack()
	logger.Error(args...)
}

func Critical(args ...interface{}) {
	debug.PrintStack()
	logger.Critical(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}