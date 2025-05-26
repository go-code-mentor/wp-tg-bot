package logger

import "log/slog"

func Debug(msg string) {
	slog.Debug(msg)
}

func Info(msg string) {
	slog.Info(msg)
}
func Warn(msg string) {
	slog.Warn(msg)
}
func Error(msg string) {
	slog.Error(msg)
}
