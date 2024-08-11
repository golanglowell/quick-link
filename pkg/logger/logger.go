package logger

import (
	"io"
	"log/slog"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger(out io.Writer) *Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &Logger{
		log: slog.New(handler),
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{log: l.log.With(args...)}
}
