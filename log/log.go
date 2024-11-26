// Package log is an abstraction layer for logging. It's a very thin
// wrapper around zerolog. In case we want to do any refactoring or
// customization later.
package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/maticnetwork/panoptichain/config"
)

func init() {
	level, err := zerolog.ParseLevel(config.Config().Logs.Verbosity)
	if err != nil {
		level = zerolog.TraceLevel
	}
	zerolog.SetGlobalLevel(level)

	if config.Config().Logs.Pretty {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return zlog.Ctx(ctx)
}

func Debug() *zerolog.Event {
	return zlog.Debug()
}

func Err(err error) *zerolog.Event {
	return zlog.Err(err)
}

func Error() *zerolog.Event {
	return zlog.Error()
}

func Fatal() *zerolog.Event {
	return zlog.Fatal()
}

func Hook(h zerolog.Hook) zerolog.Logger {
	return zlog.Hook(h)
}

func Info() *zerolog.Event {
	return zlog.Info()
}

func Level(level zerolog.Level) zerolog.Logger {
	return zlog.Level(level)
}

func Log() *zerolog.Event {
	return zlog.Log()
}

func Output(w io.Writer) zerolog.Logger {
	return zlog.Output(w)
}

func Panic() *zerolog.Event {
	return zlog.Panic()
}

func Print(v ...interface{}) {
	zlog.Print(v...)
}

func Printf(format string, v ...interface{}) {
	zlog.Printf(format, v...)
}

func Sample(s zerolog.Sampler) zerolog.Logger {
	return zlog.Sample(s)
}

func Trace() *zerolog.Event {
	return zlog.Trace()
}

func Warn() *zerolog.Event {
	return zlog.Warn()
}

func With() zerolog.Context {
	return zlog.With()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return zlog.WithLevel(level)
}
