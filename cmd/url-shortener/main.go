package main

import (
    "fmt"
    "log/slog"
    "os"
    "url-shortener/internal/config"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func main() {
    // TODO: init config: cleanenv
    cfg := config.MustLoad()

    fmt.Println(cfg)

    log := setupLogger(cfg.Env)
    log = log.With(slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

    log.Info("initializing server", slog.String("address", cfg.Address)) // Помимо сообщения выведем параметр с адресом
    log.Debug("logger debug mode enabled")

    // TODO: init logger: slog

    // TODO: init storage: sqlite

    // TODO: init router: chi, "chi render"

    // TODO: run server
}

func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env {
    case envLocal:
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envDev:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envProd:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    }

    return log
}
