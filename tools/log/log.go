package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
)

type Logger struct {
	logger *logrus.Logger
	output io.Writer
}

func New(isProd ...bool) *Logger {
	logger := logrus.New()

	level := logrus.DebugLevel
	if len(isProd) > 0 && isProd[0] == true {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@time",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, frame.File
		},
	})

	return &Logger{
		logger: logger,
	}
}

// WithField 为本次日志添加一个日志字段
func (l *Logger) WithField(key string, value interface{}) *Entry {
	entry := l.logger.WithField(key, value)
	return &Entry{entry: entry}
}

// WithFields 为本次日志添加多个日志字段
func (l *Logger) WithFields(fieldsMap map[string]interface{}) *Entry {
	entry := l.logger.WithFields(fieldsMap)
	return &Entry{entry: entry}
}

// Debug 输出Debug日志
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugf 输出Debug格式化日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info 输出Info日志
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof 输出Infof格式化日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn 输出Warn日志
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnf 输出Warnf格式化日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error 输出Error日志
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf 输出Errorf格式化日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal 输出Fatal日志
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf 输出Fatalf格式化日志
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Panic 输出Panic日志
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Panicf 输出Panicf格式化日志
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// SetOutput 设置日志输出途径
func (l *Logger) SetOutput(output io.Writer) {
	l.logger.SetOutput(output)
}

// Raw 获取logrus实例
func (l *Logger) Raw() *logrus.Logger {
	return l.logger
}
