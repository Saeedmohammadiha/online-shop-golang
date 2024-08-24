package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var logInstance Ilogger


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

type logger struct {
	Sugar   *zap.SugaredLogger
	logFile *os.File
}

func new() Ilogger {
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


	sugar := initiatedLogger.Sugar()
	return &logger{
		Sugar:   sugar,
		logFile: logFile,
	}
}

func (l *logger) Debug(msg string, args ...interface{}) {
	l.Sugar.Debugw(msg, args...)
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.Sugar.Infow(msg, args...)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.Sugar.Warnw(msg, args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.Sugar.Errorw(msg, args...)
}

func (l *logger) Fatal(msg string, args ...interface{}) {
	l.Sugar.Fatalw(msg, args...)
}

func (l *logger) Debugf(msg string, args ...interface{}) {
	l.Sugar.Debugf(msg, args...)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.Sugar.Infof(msg, args...)
}
func (l *logger) Warnf(msg string, args ...interface{}) {
	l.Sugar.Warnf(msg, args...)
}
func (l *logger) Errorf(msg string, args ...interface{}) {
	l.Sugar.Errorf(msg, args...)
}
func (l *logger) Fatalf(msg string, args ...interface{}) {
	l.Sugar.Fatalf(msg, args...)
}

func (l *logger) Sync() error {
	defer l.logFile.Close()
	return l.Sugar.Sync()
}


func init() {
	logInstance = new()
}

// Exported function to access the logger instance
func Logger() Ilogger {
	return logInstance
}