package main

import (
	"errors"
	"log"
	"log/slog"

	"github.com/tokane888/gtrading-bot/services/sample/internal/config"
	pkglogger "github.com/tokane888/trading-bot/pkg/logger"
)

// App version. Defaults to dev build. On cloud, injected at build time via -ldflags.
var version = "dev"

func main() {
	cfg, err := config.NewConfig(version)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}
	logger := pkglogger.NewLogger(cfg.Logger)

	logger.Info("sample batch info")
	logger.Info("additional field sample", slog.String("key", "value"))
	logger.Warn("sample warn")
	logger.Error("sample error")
	err = errors.New("sample error")
	logger.Error("sample error with detail", "err", err)
}
