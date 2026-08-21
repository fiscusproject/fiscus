package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fiscusproject/fiscus/internal/core"
	"github.com/fiscusproject/fiscus/internal/core/server"
	"github.com/fiscusproject/fiscus/internal/hr"
)

func main() {
	core.Initialize()

	hr.Initialize()

	serverErr := server.Start()

	exitCode := block(serverErr)

	core.Dispose()

	os.Exit(exitCode)
}

func block(serverErr <-chan error) int {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case receivedSignal := <-stop:
		slog.Info("shutdown signal received", "signal", receivedSignal.String())
		return 0
	case err := <-serverErr:
		slog.Error("server error", "error", err)
		return 1
	}
}
