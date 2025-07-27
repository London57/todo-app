package logger

import (
	"io"
	"log/slog"
)

//go:generate mockgen -source=logger.go -destination=mocks/mock.go

type Interface interface {
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
}

type Log struct {
	*slog.Logger
}

func New(writer io.Writer, opts *slog.HandlerOptions) *Log {
	return &Log{
		slog.New(
			slog.NewJSONHandler(writer, opts),
		),
	}
}

func (l *Log) Fatal(msg string, args ...any) {
	l.Error(msg, args...)
}
