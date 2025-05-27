package logger

import (
	"fmt"
	"log/slog"
)

func Debug(msg string) {
	slog.Debug(msg)
}
func Info(msg string) {
	slog.Info(msg)
}
func Infof(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}
func Warn(msg string) {
	slog.Warn(msg)
}
func Error(msg string) {
	slog.Error(msg)
}
func Errorf(format string, args ...any) {
	slog.Error(fmt.Sprintf(format, args...))
}
