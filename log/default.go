package log

import "github.com/jacksonCLyu/ridi-faces/pkg/logger"

// DefaultStdLogger is the default logger
var DefaultStdLogger logger.Logger

// Trace logger global default trace level logging function
func Trace(args ...any) {
	L().Trace(args...)
}

// Tracef logger global default trace level logging function
func Tracef(format string, args ...any) {
	L().Tracef(format, args...)
}

// Debug logger global default debug level logging function
func Debug(args ...any) {
	L().Debug(args...)
}

// Debugf logger global default debug level logging function
func Debugf(format string, args ...any) {
	L().Debugf(format, args...)
}

// Info logger global default info level logging function
func Info(args ...any) {
	L().Info(args...)
}

// Infof logger global default info level logging function
func Infof(format string, args ...any) {
	L().Infof(format, args...)
}

// Warn logger global default warn level logging function
func Warn(args ...any) {
	L().Warn(args...)
}

// Warnf logger global default warn level logging function
func Warnf(format string, args ...any) {
	L().Warnf(format, args...)
}

// Error logger global default error level logging function
func Error(args ...any) {
	L().Error(args...)
}

// Errorf logger global default error level logging function
func Errorf(format string, args ...any) {
	L().Errorf(format, args...)
}

// Fatal logger global default fatal level logging function
func Fatal(args ...any) {
	L().Fatal(args...)
}

// Fatalf logger global default fatal level logging function
func Fatalf(format string, args ...any) {
	L().Fatalf(format, args...)
}
