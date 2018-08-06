package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

func InitLog() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/tmp/log/foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	//zapcore.AddSync()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger = zap.New(core)
	sugaredLogger = logger.Sugar()
	logger.Info("logger construction succeeded")
}

// Info info
func Info(s string) {
	logger.Info(s)
}

// Infof infof
func Infof(format string, v ...interface{}) {
	sugaredLogger.Infof(format, v...)
}

// Debug debug
func Debug(s string) {
	logger.Debug(s)
}

// Debugf debugf
func Debugf(format string, v ...interface{}) {
	sugaredLogger.Debugf(format, v...)
}

// Warn warn
func Warn(s string) {
	logger.Warn(s)
}

// Warnf warnf
func Warnf(format string, v ...interface{}) {
	sugaredLogger.Warnf(format, v...)
}

// Warning warning
func Warning(s string) {
	logger.Warn(s)
}

// Warningf warningf
func Warningf(format string, v ...interface{}) {
	sugaredLogger.Warnf(format, v...)
}

// Error error
func Error(s string) {
	logger.Error(s)
}

// Errorf errorf
func Errorf(format string, v ...interface{}) {
	sugaredLogger.Errorf(format, v...)
}

// Fatal fatal
func Fatal(s string) {
	logger.Fatal(s)
}

// Fatalf fatalf
func Fatalf(format string, v ...interface{}) {
	sugaredLogger.Fatalf(format, v...)
}

// FatalEnabled fatal enabled
func FatalEnabled() bool {
	return logger.Core().Enabled(zapcore.FatalLevel)
}

// ErrorEnabled error enabled
func ErrorEnabled() bool {
	return logger.Core().Enabled(zapcore.ErrorLevel)
}

// WarnEnabled warn enabled
func WarnEnabled() bool {
	return logger.Core().Enabled(zapcore.WarnLevel)
}

// InfoEnabled info enabled
func InfoEnabled() bool {
	return logger.Core().Enabled(zapcore.InfoLevel)
}

// DebugEnabled debug enabled
func DebugEnabled() bool {
	return logger.Core().Enabled(zapcore.DebugLevel)
}
