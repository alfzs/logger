# Logger Package

Пакет logger предоставляет удобную обертку для работы с логгером `slog` с поддержкой мультиплексирования вывода (консоль + файл) и ротации логов.

## Установка

```bash
go get github.com/alfzs/logger
```

````

## Использование

### Быстрый старт

```go
import "github.com/alfzs/logger"

// Реализуйте интерфейс LoggerConfig
type AppConfig struct {
    LogLevel      string
    LogFile       string
    LogMaxSize    int
    LogMaxBackups int
    LogMaxAge     int
    LogCompress   bool
}

func (c AppConfig) GetLevel() string      { return c.LogLevel }
func (c AppConfig) GetFile() string       { return c.LogFile }
func (c AppConfig) GetMaxSizeMB() int     { return c.LogMaxSize }
func (c AppConfig) GetMaxBackups() int    { return c.LogMaxBackups }
func (c AppConfig) GetMaxAgeDays() int    { return c.LogMaxAge }
func (c AppConfig) GetCompress() bool     { return c.LogCompress }

func main() {
    cfg := AppConfig{
        LogLevel:      "debug",
        LogFile:       "app.log",
        LogMaxSize:    100, // MB
        LogMaxBackups: 3,
        LogMaxAge:     30, // days
        LogCompress:   true,
    }

    log := logger.NewLogger(logger.LoggerParams{Configs: cfg})

    log.Info("Application started")
    log.Debug("Debug information")
    log.Error("Error message", slog.String("error", err.Error()))
}
```

### Особенности

- **Мульти-хендлер**: логи пишутся одновременно в консоль (текстовый формат) и в файл (JSON формат)
- **Ротация логов**: автоматическая ротация логов с настройками:
  - Максимальный размер файла
  - Максимальное количество бэкапов
  - Максимальный возраст файлов
  - Сжатие старых логов
- **Уровни логирования**: debug, info, warn, error

## Конфигурация

Реализуйте интерфейс `LoggerConfig` в вашей структуре конфигурации:

```go
type LoggerConfig interface {
    GetLevel() string     // Уровень логирования: debug, info, warn, error
    GetFile() string      // Путь к файлу логов
    GetMaxSizeMB() int    // Макс. размер файла в MB перед ротацией
    GetMaxBackups() int   // Макс. количество старых лог-файлов
    GetMaxAgeDays() int   // Макс. возраст файлов в днях
    GetCompress() bool    // Сжимать ли старые логи
}
```

## Пример вывода

**Консоль (текстовый формат):**

```
time=2023-11-15T12:00:00.000Z level=INFO msg="Application started"
```

**Файл (JSON формат):**

```json
{ "time": "2023-11-15T12:00:00.000Z", "level": "INFO", "msg": "Application started" }
```

## Зависимости

- [slog](https://pkg.go.dev/log/slog) - стандартный логгер Go
- [lumberjack](https://github.com/natefinch/lumberjack) - ротация лог-файлов

## Лицензия

MIT
````
