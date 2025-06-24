package logger

import (
	"context"
	"log/slog"
)

// MultiHandler реализует slog.Handler для мультиплексирования логов между несколькими обработчиками
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler создает новый мульти-хендлер
func NewMultiHandler(handlers ...slog.Handler) slog.Handler {
	return &MultiHandler{handlers: handlers}
}

// Enabled проверяет, активен ли хотя бы один обработчик для данного уровня
func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle передает запись всем обработчикам
func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, record)
	}
	return nil
}

// WithAttrs создает новый хендлер с добавленными атрибутами
func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var handlers []slog.Handler
	for _, h := range m.handlers {
		handlers = append(handlers, h.WithAttrs(attrs))
	}
	return NewMultiHandler(handlers...)
}

// WithGroup создает новый хендлер с добавленной группой
func (m *MultiHandler) WithGroup(name string) slog.Handler {
	var handlers []slog.Handler
	for _, h := range m.handlers {
		handlers = append(handlers, h.WithGroup(name))
	}
	return NewMultiHandler(handlers...)
}
