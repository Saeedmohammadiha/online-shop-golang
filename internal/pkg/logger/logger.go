package logger

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	Sugar   *zap.SugaredLogger
	logFile *os.File
}

func New() Ilogger {
	logFile, err := os.Create("logfile.log")
	if err != nil {
		panic(err)
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // or NewConsoleEncoder
		zapcore.AddSync(logFile),
		zap.InfoLevel, // Log level
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	initiatedLogger := zap.New(zapcore.NewTee(fileCore, consoleCore))
	// Create a logger from the file core
	//initiatedLogger, err := zap.NewProduction()

	if err != nil {
		slog.Error("logger app could not be initiate", "initialZapError:", err.Error())
	}

	sugar := initiatedLogger.Sugar()
	return &Logger{
		Sugar:   sugar,
		logFile: logFile,
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
	defer l.logFile.Close()
	return l.Sugar.Sync()
}
