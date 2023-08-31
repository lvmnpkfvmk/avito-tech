package main

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lvmnpkfvmk/avito-tech/config"
	"github.com/lvmnpkfvmk/avito-tech/internal/logger"
	"github.com/lvmnpkfvmk/avito-tech/internal/handlers"
	"github.com/lvmnpkfvmk/avito-tech/internal/repository"
)

var repo repository.SegmentRepository

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


	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	sHandler := handlers.NewSegmentHandler(repo, logger)

	segmentRoute := e.Group("/segment")
	segmentRoute.POST("/", sHandler.CreateSegment)
	segmentRoute.DELETE("/", sHandler.DeleteSegment)

	e.Logger.Fatal(e.Start(cfg.HTTPAddr))
	return nil
}
