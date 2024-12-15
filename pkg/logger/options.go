package logger

import (
	"io"
	"os"
	"path/filepath"

	"gitee.com/qciip-icp/v-trace/pkg/tools/pathtools"
)

const (
	// FmtEmptySeparate represents empty format string.
	FmtEmptySeparate = ""
)

// Level is the log level.
type Level uint8

const (
	// DebugLevel only use when dev.
	DebugLevel Level = iota

	// InfoLevel log some useful info.
	InfoLevel

	// WarnLevel is more important than InfoLevel.
	WarnLevel

	// ErrorLevel indicate that the program had an error occurred.
	ErrorLevel

	// PanicLevel log is print when a panic occurred.
	PanicLevel

	// FatalLevel log is print when program can not run anymore.
	FatalLevel
)

// LevelNameMapping is a mapping from log level to level name.
var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

type Config struct {
	Level             int     `mapstructure:"level"`
	Formatter         string  `mapstructure:"formatter"`
	DisableCaller     *bool   `mapstructure:"disable_caller"`
	DisableColor      *bool   `mapstructure:"disable_color"`
	EnableAbsPath     *bool   `mapstructure:"enable_abs_path"`
	EnableFunc        *bool   `mapstructure:"enable_func"`
	IgnoreBasicFields *bool   `mapstructure:"ignore_basic_fields"`
	OutputFile        *string `mapstructure:"output_file"`

	// below only support on text format now
	Prefix *string           `mapstructure:"prefix"`
	Kv     map[string]string `mapstructure:"kv"`
}

func setLogger(c *Config, l *Logger) {
	var formatter Formatter
	var ignoreBasicFields bool

	if c.IgnoreBasicFields == nil {
		ignoreBasicFields = true
	} else {
		ignoreBasicFields = false
	}

	switch c.Formatter {
	case "text":
		formatter = &TextFormatter{
			ignoreBasicFields,
		}
	case "json":
		formatter = &JsonFormatter{
			ignoreBasicFields,
		}
	default:
		formatter = &TextFormatter{
			ignoreBasicFields,
		}
	}

	setOptions(
		l,
		WithLevel(Level(c.Level)),
		WithFormatter(formatter),
		WithDisableCaller(c.DisableCaller),
		WithDisableColor(c.DisableColor),
		WithEnableAbsPath(c.EnableAbsPath),
		WithEnableFunc(c.EnableFunc),
		WithPrefix(c.Prefix),
		WithKv(c.Kv),
	)

	if c.OutputFile != nil && *c.OutputFile != "" {
		pathtools.MkdirIfNotExist(filepath.Dir(*c.OutputFile))
		f, err := os.OpenFile(*c.OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {

			return
		}
		setOptions(
			l,
			WithOutput(f),
		)
	} else {
		setOptions(
			l,
			WithOutput(os.Stdout),
		)
	}
}

func SetLogger(c *Config, l *Logger) {
	if l == nil {
		setLogger(c, std)
		return
	}

	setLogger(c, l)
}

// log options.
type options struct {
	output    io.Writer
	level     Level
	stdLevel  Level
	formatter Formatter

	// if true, will not print call for log
	disableCaller bool
	// if true, will print log without color
	disableColor bool
	// if true, will print the absolute path about the file which print the log, default false
	enableAbsPath bool
	// if true will print the func which do log, default false
	enableFunc bool

	prefix string

	kv map[string]string
}

type Option func(*options)

func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

// WithOutput set log output.
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithLevel sets log level.
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// WithStdLevel sets log std level.
func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

// WithFormatter sets log formatter.
func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

// WithDisableCaller sets log disable caller. if caller is true, then the log location will not be printed.
func WithDisableCaller(disableCaller *bool) Option {
	return func(o *options) {
		if disableCaller != nil {
			o.disableCaller = *disableCaller
			return
		}

		o.disableCaller = true
	}
}

// WithDisableColor sets log color. if disableColor is true, then the log color will not be printed.
func WithDisableColor(disableColor *bool) Option {
	return func(o *options) {
		if disableColor != nil {
			o.disableColor = *disableColor
			return
		}

		o.disableColor = false
	}
}

func WithEnableAbsPath(enableAbsPath *bool) Option {
	return func(o *options) {
		if enableAbsPath != nil {
			o.enableAbsPath = *enableAbsPath
			return
		}

		o.enableAbsPath = false
	}
}

func WithEnableFunc(enableFunc *bool) Option {
	return func(o *options) {
		if enableFunc != nil {
			o.enableFunc = *enableFunc
			return
		}

		o.enableFunc = false
	}
}

func WithPrefix(prefix *string) Option {
	return func(o *options) {
		if prefix != nil {
			o.prefix = *prefix
			return
		}

		o.prefix = ""
	}
}

func WithKv(kv map[string]string) Option {
	return func(o *options) {
		if kv != nil {
			o.kv = kv

			return
		}

		o.kv = make(map[string]string)
	}
}
