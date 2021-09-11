package log

import "github.com/sirupsen/logrus"

type Entry struct {
	entry *logrus.Entry
}

// Debug 输出Debug日志
func (e *Entry) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

// Debugf 输出Debug格式化日志
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

// Info 输出Info日志
func (e *Entry) Info(args ...interface{}) {
	e.entry.Info(args...)
}

// Infof 输出Infof格式化日志
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Warn 输出Warn日志
func (e *Entry) Warn(args ...interface{}) {
	e.entry.Warn(args...)
}

// Warnf 输出Warnf格式化日志
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

// Error 输出Error日志
func (e *Entry) Error(args ...interface{}) {
	e.entry.Error(args...)
}

// Errorf 输出Errorf格式化日志
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Fatal 输出Fatal日志
func (e *Entry) Fatal(args ...interface{}) {
	e.entry.Fatal(args...)
}

// Fatalf 输出Fatalf格式化日志
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

// Panic 输出Panic日志
func (e *Entry) Panic(args ...interface{}) {
	e.entry.Panic(args...)
}

// Panicf 输出Panicf格式化日志
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.entry.Panicf(format, args...)
}

// Raw 获取logrus.Entry实例
func (e *Entry) Raw() *logrus.Entry {
	return e.entry
}
