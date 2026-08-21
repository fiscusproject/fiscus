package core

import (
	"context"
	"log/slog"
	"time"

	"github.com/fiscusproject/fiscus/internal/core/commons"
	"github.com/fiscusproject/fiscus/internal/core/internal/api"
	"github.com/fiscusproject/fiscus/internal/core/internal/environment"
	"github.com/fiscusproject/fiscus/internal/core/internal/logging"
	"github.com/fiscusproject/fiscus/internal/core/internal/telemetry"
	"github.com/fiscusproject/fiscus/internal/core/server"
)

func Initialize() {
	environment.Load()
	logging.Initialize()
	slog.Info("starting", "service", commons.ServiceName, "version", commons.Version)

	telemetry.Initialize()

	server.Initialize()
	api.RegisterRoutes()
}

func Dispose() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	telemetry.Dispose(ctx)
}
