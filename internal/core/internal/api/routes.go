package api

import (
	"github.com/fiscusproject/fiscus/internal/core/server"
)

func RegisterRoutes() {
	server.GetRouter().GET("/healthz", Healthz)
	server.GetRouter().GET("/readyz", Readyz)
}
