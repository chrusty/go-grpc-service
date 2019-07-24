package interfaces

// Logger is a simplified logging interface:
type Logger interface {
	Tracef(message string, args ...interface{})
	Debugf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
}
