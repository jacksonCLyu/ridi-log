package log

import (
	"io"

	"github.com/jacksonCLyu/ridi-faces/pkg/configer"
	"github.com/jacksonCLyu/ridi-faces/pkg/logger"
)

type InitOption interface {
	initApply(opts *initOptions)
}

type initOptions struct {
	logger logger.Logger
	config configer.Configurable
}

type initOptionFunc func(*initOptions)

func (f initOptionFunc) initApply(opts *initOptions) {
	f(opts)
}

// WithLogger sets the logger to be used by the logger.
func WithLogger(logger logger.Logger) InitOption {
	return initOptionFunc(func(opts *initOptions) {
		opts.logger = logger
	})
}

// WithConfig sets the global config.
func WithConfig(config configer.Configurable) InitOption {
	return initOptionFunc(func(opts *initOptions) {
		opts.config = config
	})
}

type Option interface {
	apply(*options)
}

type applyFunc func(*options)

func (f applyFunc) apply(o *options) {
	f(o)
}

type options struct {
	name       string
	category   string
	fields     map[string]any
	caller     bool
	callerSkip int
	writers    []io.Writer
}

// WithName sets the name of the logger.
func WithName(name string) Option {
	return applyFunc(func(o *options) {
		o.name = name
	})
}

// WithCategory sets the category of the logger.
func WithCategory(category string) Option {
	return applyFunc(func(o *options) {
		o.category = category
	})
}

// AddField adds a field to the logger.
func AddField(key string, value any) Option {
	return applyFunc(func(o *options) {
		if o.fields == nil {
			o.fields = make(map[string]any)
		}
		o.fields[key] = value
	})
}

// Caller adds the caller to the log message.
func Caller() Option {
	return applyFunc(func(o *options) {
		o.caller = true
	})
}

// WithCaller sets the caller.
func WithCaller(caller bool) Option {
	return applyFunc(func(o *options) {
		o.caller = caller
	})
}

// CallerSkip sets the caller skip.
func CallerSkip(skip int) Option {
	return applyFunc(func(o *options) {
		o.callerSkip = skip
	})
}

// AddWriter sets the log writer.
func AddWriter(writer io.Writer) Option {
	return applyFunc(func(o *options) {
		o.writers = append(o.writers, writer)
	})
}
