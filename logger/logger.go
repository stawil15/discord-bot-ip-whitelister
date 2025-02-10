package logger

import (
	"log/slog"
	"os"

	"github.com/geekloper/discord-bot-ip-whitelister/config"
)

// Global logger instance
var Log *slog.Logger

func InitLogger() {
	level := slog.LevelInfo

	if config.DebugMode() {
		level = slog.LevelDebug
	}

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	Log = slog.New(logHandler)
}

// Debug logs
func Debug(msg string, keysAndValues ...interface{}) {
	Log.Debug(msg, keysAndValues...)
}

// Info logs
func Info(msg string, keysAndValues ...interface{}) {
	Log.Info(msg, keysAndValues...)
}

// Warning logs
func Warn(msg string, keysAndValues ...interface{}) {
	Log.Warn(msg, keysAndValues...)
}

// Error logs
func Error(msg string, keysAndValues ...interface{}) {
	Log.Error(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	Log.Error(msg, keysAndValues...)
	os.Exit(1) // Terminate application
}
