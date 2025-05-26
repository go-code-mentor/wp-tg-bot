package logger

import "log/slog"

func Debug(msg string, args ...any) {
	slog.Debug(msg, args)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args)
}
func Warn(msg string, args ...any) {
	slog.Warn(msg, args)
}
func Error(msg string, args ...any) {
	slog.Error(msg, args)
}
