package logger

// Logger представляет собой интерфейс логгера
type Logger interface {
	Init()

	Debug(args ...any)
	Debugf(format string, args ...any)

	Info(args ...any)
	Infof(format string, args ...any)

	Warn(args ...any)
	Warnf(format string, args ...any)

	Error(args ...any)
	Errorf(format string, args ...any)

	Panic(args ...any)
	Panicf(format string, args ...any)

	Fatal(args ...any)
	Fatalf(format string, args ...any)

	Print(args ...any)
	Printf(format string, args ...any)
}
