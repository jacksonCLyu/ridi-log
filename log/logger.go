package log

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/jacksonCLyu/ridi-config/pkg/config"
	"github.com/jacksonCLyu/ridi-faces/pkg/configer"
	"github.com/jacksonCLyu/ridi-faces/pkg/logger"
	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/objects"
)

type simpleLogger struct {
	stdLogger *log.Logger
	opts      *options
}

var globalConfig configer.Configurable

func Init(opts ...InitOption) error {
	initOpts := &initOptions{
		logger: nil,
		config: config.L(),
	}
	for _, opt := range opts {
		opt.initApply(initOpts)
	}
	if initOpts.logger != nil {
		DefaultStdLogger = initOpts.logger
		return nil
	}
	DefaultStdLogger = NewLogger()
	return nil
}

// SetGlobalConfig sets the global config.
func SetGlobalConfig(c configer.Configurable) {
	globalConfig = c
}

var _ logger.Logger = (*simpleLogger)(nil)

// NewLogger returns a new logger. default writer is os.Stdout
func NewLogger(opts ...Option) logger.Logger {
	// default options
	options := &options{
		name:       "defaultLog",
		category:   "app",
		caller:     true,
		callerSkip: 2,
		writers:    []io.Writer{},
	}
	for _, o := range opts {
		o.apply(options)
	}
	b := bytes.NewBuffer(nil)
	b.WriteString(" [")
	b.WriteString(options.category)
	b.WriteString("] ")
	b.WriteString("[")
	b.WriteString(options.name)
	b.WriteString("] - ")
	return &simpleLogger{
		stdLogger: log.New(os.Stdout, b.String(), log.LstdFlags),
		opts:      options,
	}
}

// L returns the global default logger.
func L() logger.Logger {
	if DefaultStdLogger == nil {
		_ = Init()
	}
	return DefaultStdLogger
}

func (l simpleLogger) Trace(args ...any) {
	l.Println(args...)
}

func (l simpleLogger) Tracef(format string, args ...any) {
	l.Printf(format, args...)
}

func (l simpleLogger) Debug(args ...any) {
	if !l.isEnable(logger.LogLevelDebug) {
		return
	}
	l.Println(args...)
}

func (l simpleLogger) Debugf(format string, args ...any) {
	if !l.isEnable(logger.LogLevelDebug) {
		return
	}
	l.Printf(format, args...)
}

func (l simpleLogger) Info(args ...any) {
	if !l.isEnable(logger.LogLevelInfo) {
		return
	}
	l.Println(args...)
}

func (l simpleLogger) Infof(format string, args ...any) {
	if !l.isEnable(logger.LogLevelInfo) {
		return
	}
	l.Printf(format, args...)
}

func (l simpleLogger) Warn(args ...any) {
	if !l.isEnable(logger.LogLevelWarn) {
		return
	}
	l.Println(args...)
}

func (l simpleLogger) Warnf(format string, args ...any) {
	if !l.isEnable(logger.LogLevelWarn) {
		return
	}
	l.Printf(format, args...)
}

func (l simpleLogger) Error(args ...any) {
	if !l.isEnable(logger.LogLevelError) {
		return
	}
	l.Println(args...)
}

func (l simpleLogger) Errorf(format string, args ...any) {
	if !l.isEnable(logger.LogLevelError) {
		return
	}
	l.Printf(format, args...)
}

func (l simpleLogger) Fatal(args ...any) {
	if !l.isEnable(logger.LogLevelFatal) {
		return
	}
	l.Println(args...)
}

func (l simpleLogger) Fatalf(format string, args ...any) {
	if !l.isEnable(logger.LogLevelFatal) {
		return
	}
	l.Printf(format, args...)
}

func (l simpleLogger) isEnable(level logger.LogLevel) bool {
	objects.RequireNonNil(globalConfig)
	configLogLevel := assignutil.Assign((globalConfig.GetInt("log.level")))
	return int(level) >= configLogLevel
}

func (l simpleLogger) Println(args ...any) {
	l.stdLogger.Println(args...)
}

func (l simpleLogger) Printf(format string, args ...any) {
	l.stdLogger.Printf(format, args...)
}
