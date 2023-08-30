package main

import (
	"context"
	"log/slog"

	"github.com/lvmnpkfvmk/avito-tech/config"
	"github.com/lvmnpkfvmk/avito-tech/internal/logger"
	"github.com/lvmnpkfvmk/avito-tech/internal/repository"
)

func main() {
	cfg := config.Get()
	ctx := context.Background()

	logger := logger.SetupLogger(cfg.LogLevel)
	if err := run(logger, cfg, ctx); err != nil {
		logger.Error("Error", err)
	}
}

func run(logger *slog.Logger, cfg *config.Config, ctx context.Context) error {
	repo, err := repository.NewSegmentRepository(ctx, cfg)
	if err != nil {
		logger.Error("Error creating repository", err)
	}
	logger.Debug("Repository is ready", repo)


	return nil
}
