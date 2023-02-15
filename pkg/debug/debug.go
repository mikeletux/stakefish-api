package debug

import (
	"log"
	"os"
)

// Logger is the interface that needs to be implemented to debug.
type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

// BuiltinLogger is a simple wrapper to debug.Logger.
type BuiltinLogger struct {
	logger *log.Logger
}

// NewBuiltinStdoutLogger returns a simple BuiltinLogger that only logs on stdout.
func NewBuiltinStdoutLogger() *BuiltinLogger {
	return &BuiltinLogger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (b *BuiltinLogger) Debug(args ...any) {
	b.logger.Println(args...)
}
func (b *BuiltinLogger) Debugf(format string, args ...any) {
	b.logger.Printf(format, args...)
}
func (b *BuiltinLogger) Fatal(args ...any) {
	b.logger.Fatalln(args...)
}
func (b *BuiltinLogger) Fatalf(format string, args ...any) {
	b.logger.Fatalf(format, args...)
}
