package logger

import (
	"log/slog"

	"go.uber.org/zap"
)

type Ilogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Sync() error
}

type Logger struct {
	Sugar *zap.SugaredLogger
}

func New() Ilogger {

	initiatedLogger, err := zap.NewProduction()

	if err != nil {
		slog.Error("logger app could not be initiate", "initialZapError:", err.Error())
	}

	sugar := initiatedLogger.Sugar()
	return &Logger{
		Sugar: sugar,
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Sugar.Debugw(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Sugar.Infow(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Sugar.Warnw(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Sugar.Errorw(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Sugar.Fatalw(msg, args...)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.Sugar.Debugf(msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.Sugar.Infof(msg, args...)
}
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.Sugar.Warnf(msg, args...)
}
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.Sugar.Errorf(msg, args...)
}
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.Sugar.Fatalf(msg, args...)
}

func (l *Logger) Sync() error {
	return l.Sugar.Sync()
}
