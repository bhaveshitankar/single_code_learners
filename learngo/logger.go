package learngo

import (
	"log/slog"
	"os"
	"sync"
)

func SetupDefaultLogger() *slog.Logger {
	// logHandeller := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level:     slog.LevelDebug,
	// 	AddSource: true,
	// }) info-->debug-->warning-->error
	logHandeller := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: true,
	})
	logger := slog.New(logHandeller)
	slog.SetDefault(logger)
	return logger
}

var (
	instance *slog.Logger
	once     sync.Once
)

func GetSingleInstanceLogger() *slog.Logger {
	once.Do(func() {
		instance = SetupDefaultLogger()
	})
	return instance
}
