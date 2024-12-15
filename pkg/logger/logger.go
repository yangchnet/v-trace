package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

var (
	std     *Logger
	stdOnce = sync.Once{}
)

type Logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func init() {
	stdOnce.Do(func() {
		logger := New()

		std = logger
	})
}

func New(opts ...Option) *Logger {
	l := &Logger{
		opt: &options{
			output:   os.Stdout,
			level:    0,
			stdLevel: 0,
			formatter: &TextFormatter{
				true,
			},
			disableCaller: false,
			disableColor:  false,
			enableAbsPath: false,
			enableFunc:    false,
			prefix:        "",
			kv:            map[string]string{},
		},
	}
	l.entryPool = &sync.Pool{New: func() interface{} { return entry(l) }}

	for _, op := range opts {
		op(l.opt)
	}

	return l
}

func StdLogger() *Logger {
	return std
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func setOptions(l *Logger, opts ...Option) {
	l.SetOptions(opts...)
}

func (l *Logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func Writer() io.Writer {
	return std
}

func (l *Logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *Logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

// std Logger.
func Debug(args ...interface{}) {
	std.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	std.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	std.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	std.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	std.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	std.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	std.entry().write(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	std.entry().write(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.entry().write(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.entry().write(ErrorLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	std.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}
