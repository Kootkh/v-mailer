/*
https://betterstack.com/community/guides/logging/logging-best-practices/
https://betterstack.com/community/guides/logging/log-formatting/
https://betterstack.com/community/guides/logging/logging-in-go/
https://betterstack.com/community/guides/logging/json-logging/
Structured Logging with slog - Part 2 - https://www.youtube.com/watch?v=ReWtK-HUbQQ
*/
package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// -----------------------------------------------------------

const (
	defaultLevel      = LevelInfo
	defaultAddSource  = true
	defaultIsJSON     = true
	defaultSetDefault = true
)

const (
	defLogLevel   = LevelInfo
	defAddSource  = true
	defLogFormat  = "json"
	defSetDefault = true
)

var defLogDest = os.Stdout

type Format string
type LogDest io.Writer

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

// -----------------------------------------------------------

func NewLogger(opts ...LoggerOption) *Logger {
	config := &LoggerOptions{
		Level:      defaultLevel,
		AddSource:  defaultAddSource,
		IsJSON:     defaultIsJSON,
		SetDefault: defaultSetDefault,
	}

	for _, opt := range opts {
		opt(config)
	}

	options := &HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	var h Handler = NewTextHandler(os.Stdout, options)

	if config.IsJSON {
		h = NewJSONHandler(os.Stdout, options)
	}

	logger := New(h)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

func NewLLogger(opts ...LLogOption) *Logger {
	config := &LLogOptions{
		Level:      defLogLevel,
		AddSource:  defAddSource,
		Format:     defLogFormat,
		SetDefault: defSetDefault,
		LogDest:    defLogDest,
	}

	for _, opt := range opts {
		opt(config)
	}

	options := &HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	//var h Handler = NewTextHandler(os.Stdout, options)
	//var h Handler = NewJSONHandler(os.Stdout, options)
	var h Handler = NewTextHandler(os.Stdout, options)

	if config.Format == "json" {
		switch config.LogDest {
		case nil:
			h = NewJSONHandler(os.Stdout, options)
		default:
			h = NewJSONHandler(config.LogDest, options)
		}
	} else if config.Format == "text" {
		switch config.LogDest {
		case nil:
			h = NewTextHandler(os.Stdout, options)
		default:
			h = NewTextHandler(config.LogDest, options)
		}
	}

	logger := New(h)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

// -----------------------------------------------------------

type LoggerOptions struct {
	Level      Level
	AddSource  bool
	IsJSON     bool
	SetDefault bool
}

type LoggerOption func(*LoggerOptions)

type LLogOptions struct {
	Level      Level
	AddSource  bool
	Format     string
	SetDefault bool
	LogDest    io.Writer
}

type LLogOption func(*LLogOptions)

// -----------------------------------------------------------

// WithLevel logger option sets the log level, if not set, the default level is Info.
func WithLevel(level string) LoggerOption {
	return func(o *LoggerOptions) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		o.Level = l
	}
}

// WithLLevel logger option sets the log level, if not set, the default level is Info.
func WithLLevel(level string) LLogOption {
	return func(o *LLogOptions) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		o.Level = l
	}
}

// -----------------------------------------------------------

// WithAddSource logger option sets the add source option, which will add source file and line number to the log record.
func WithAddSource(addSource bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.AddSource = addSource
	}
}

// WithLAddSource logger option sets the add source option, which will add source file and line number to the log record.
func WithLAddSource(addSource bool) LLogOption {
	return func(o *LLogOptions) {
		o.AddSource = addSource
	}
}

// -----------------------------------------------------------

// WithIsJSON logger option sets the is json option, which will set JSON format for the log record.
func WithIsJSON(isJSON bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.IsJSON = isJSON
	}
}

func WithFormat(format string) LLogOption {
	return func(o *LLogOptions) {
		var f Format
		switch format {
		case "json":
			f = FormatJSON
		case "text":
			f = FormatText
		default:
			f = FormatJSON
		}

		o.Format = string(f)
	}
}

func WithLogDest(ld io.Writer) LLogOption {
	return func(o *LLogOptions) {
		o.LogDest = ld
	}
}

// -----------------------------------------------------------

// WithSetDefault logger option sets the set default option, which will set the created logger as default logger.
func WithSetDefault(setDefault bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.SetDefault = setDefault
	}
}

// WithSetDefault logger option sets the set default option, which will set the created logger as default logger.
func WithLSetDefault(setDefault bool) LLogOption {
	return func(o *LLogOptions) {
		o.SetDefault = setDefault
	}
}

// -----------------------------------------------------------

// WithAttrs returns logger with attributes.
func WithAttrs(ctx context.Context, attrs ...Attr) *Logger {
	logger := L(ctx)
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

// -----------------------------------------------------------

// WithDefaultAttrs returns logger with default attributes.
func WithDefaultAttrs(logger *Logger, attrs ...Attr) *Logger {
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

// -----------------------------------------------------------

func L(ctx context.Context) *Logger {
	return loggerFromContext(ctx)
}

/* func LL(ctx context.Context, l *Logger) *Logger {
	return LFromContext(ctx, l)
} */

// -----------------------------------------------------------

func Default() *Logger {
	return slog.Default()
}
