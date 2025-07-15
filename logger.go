package logger

import (
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

// Config хранит настройки Logger
type Config struct {
	Level      string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
	File       string `yaml:"file" env:"LOG_FILE" env-default:"logs/app.log"`
	MaxSizeMB  int    `yaml:"max_size_mb" env:"LOG_MAX_SIZE_MB" env-default:"64"`
	MaxBackups int    `yaml:"max_backups" env:"LOG_MAX_BACKUPS" env-default:"7"`
	MaxAgeDays int    `yaml:"max_age_days" env:"LOG_MAX_AGE_DAYS" env-default:"30"`
	Compress   bool   `yaml:"compress" env:"LOG_COMPRESS" env-default:"true"`
}

// NewLogger создает новый экземпляр логгера с мульти-хендлером (консоль + файл) на основе переданной конфигурации
func NewLogger(c Config) *slog.Logger {
	var level slog.Level
	switch c.Level {
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
		Filename:   c.File,
		MaxSize:    c.MaxSizeMB,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAgeDays,
		Compress:   c.Compress,
	}

	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	fileHandler := slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})

	return slog.New(NewMultiHandler(consoleHandler, fileHandler))
}
