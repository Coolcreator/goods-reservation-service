package logger

var _ Logger = (*loggerStub)(nil)

type loggerStub struct{}

func NewLoggerStub() *loggerStub {
	return &loggerStub{}
}

func (l *loggerStub) Init() {}

func (l *loggerStub) Debug(args ...any) {}

func (l *loggerStub) Debugf(template string, args ...any) {}

func (l *loggerStub) Info(args ...any) {}

func (l *loggerStub) Infof(template string, args ...any) {}

func (l *loggerStub) Warn(args ...any) {}

func (l *loggerStub) Warnf(template string, args ...any) {}

func (l *loggerStub) Error(args ...any) {}

func (l *loggerStub) Errorf(template string, args ...any) {}

func (l *loggerStub) Panic(args ...any) {}

func (l *loggerStub) Panicf(template string, args ...any) {}

func (l *loggerStub) Fatal(args ...any) {}

func (l *loggerStub) Fatalf(template string, args ...any) {}

func (l *loggerStub) Print(args ...any) {}

func (l *loggerStub) Printf(template string, args ...any) {}
