package logger

import (
	"context"
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	zap *zap.Logger
}

func NewLogger() handler.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "datetime"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	l, err := config.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &logger{
		zap: l,
	}
}

// Info logs an informational message.
// ctx: context for the log entry
// msg: the message to log
func (l *logger) Info(ctx context.Context, msg string) {
	l.zap.Info(msg)
}

// Warn logs a warn message.
// ctx: context for the log entry
// msg: the message to log
func (l *logger) Warn(ctx context.Context, msg string) {
	l.zap.Warn(msg)
}

// Error logs an error message with an error.
// ctx: context for the log entry
// msg: the message to log
// err: the error to log
func (l *logger) Error(ctx context.Context, msg string, err error) {
	l.zap.Error(msg, zap.Error(err))
}

// Panic logs a panic message with an error.
// ctx: context for the log entry
// msg: the message to log
func (l *logger) Panic(ctx context.Context, msg string) {
	l.zap.Panic(msg)
}
