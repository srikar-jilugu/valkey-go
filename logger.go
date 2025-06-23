package valkey

import "log"

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type defaultLogger struct{}

func (l *defaultLogger) Debugf(format string, args ...interface{}) {
	// no ops
}
func (l *defaultLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}
func (l *defaultLogger) Warnf(format string, args ...interface{}) {
	// no ops
}
func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
func (l *defaultLogger) Fatalf(format string, args ...interface{}) {
	// no ops
}

var globalLogger Logger = &defaultLogger{}

func SetGlobalLogger(logger Logger) {
	if logger == nil {
		globalLogger = &defaultLogger{}
	} else {
		globalLogger = logger
	}
}
