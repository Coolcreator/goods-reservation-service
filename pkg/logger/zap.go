package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*zapSugaredLogger)(nil)

// zapSugaredLogger является оберткой над zap.SugaredLogger,
// имплементирует интерфейс логгера
type zapSugaredLogger struct {
	cfg   *Config
	sugar *zap.SugaredLogger
}

// Config представляет собой конфигурацию zapSugaredLogger
type Config struct {
	Development bool
	Level       string
	Encoding    string
}

// NewZapSugaredlogger создает новый экземпляр zapSugaredLogger
func NewZapSugaredlogger(cfg *Config) *zapSugaredLogger {
	return &zapSugaredLogger{cfg: cfg}
}

// маппинг уровня логирования из конфигурации в zapcore.Level
var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

// level определяет уровень логирования
func (l *zapSugaredLogger) level() zapcore.Level {
	level, exist := loggerLevelMap[l.cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init инициализирует zapSugaredLogger
func (l *zapSugaredLogger) Init() {
	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.StampMilli)

	var encoder zapcore.Encoder
	if l.cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(l.level()))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugar = logger.Sugar()
}

func (l *zapSugaredLogger) Debug(args ...any) {
	l.sugar.Debug(args...)
}

func (l *zapSugaredLogger) Debugf(template string, args ...any) {
	l.sugar.Debugf(template, args...)
}

func (l *zapSugaredLogger) Info(args ...any) {
	l.sugar.Info(args...)
}

func (l *zapSugaredLogger) Infof(template string, args ...any) {
	l.sugar.Infof(template, args...)
}

func (l *zapSugaredLogger) Warn(args ...any) {
	l.sugar.Warn(args...)
}

func (l *zapSugaredLogger) Warnf(template string, args ...any) {
	l.sugar.Warnf(template, args...)
}

func (l *zapSugaredLogger) Error(args ...any) {
	l.sugar.Error(args...)
}

func (l *zapSugaredLogger) Errorf(template string, args ...any) {
	l.sugar.Errorf(template, args...)
}

func (l *zapSugaredLogger) Panic(args ...any) {
	l.sugar.Panic(args...)
}

func (l *zapSugaredLogger) Panicf(template string, args ...any) {
	l.sugar.Panicf(template, args...)
}

func (l *zapSugaredLogger) Fatal(args ...any) {
	l.sugar.Fatal(args...)
}

func (l *zapSugaredLogger) Fatalf(template string, args ...any) {
	l.sugar.Fatalf(template, args...)
}

func (l *zapSugaredLogger) Print(args ...any) {
	l.sugar.Info(args...)
}

func (l *zapSugaredLogger) Printf(template string, args ...any) {
	l.sugar.Infof(template, args...)
}
