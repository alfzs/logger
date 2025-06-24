package logger

import (
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

// LoggerConfig интерфейс конфигурации логгера
type LoggerConfig interface {
	GetLevel() string
	GetFile() string
	GetMaxSizeMB() int
	GetMaxBackups() int
	GetMaxAgeDays() int
	GetCompress() bool
}

// LoggerParams параметры для создания логгера
type LoggerParams struct {
	Configs LoggerConfig
}

// NewLogger создает новый экземпляр логгера с мульти-хендлером (консоль + файл) на основе переданной конфигурации
func NewLogger(p LoggerParams) *slog.Logger {
	var level slog.Level
	switch p.Configs.GetLevel() {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	writer := &lumberjack.Logger{
		Filename:   p.Configs.GetFile(),
		MaxSize:    p.Configs.GetMaxSizeMB(),
		MaxBackups: p.Configs.GetMaxBackups(),
		MaxAge:     p.Configs.GetMaxAgeDays(),
		Compress:   p.Configs.GetCompress(),
	}

	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	fileHandler := slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})

	return slog.New(NewMultiHandler(consoleHandler, fileHandler))
}
