package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/fiscusproject/fiscus/internal/core/commons"
	"github.com/fiscusproject/fiscus/internal/core/internal/environment"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server struct {
	router       *gin.Engine
	fiscusRouter *gin.RouterGroup
	httpServer   *http.Server
}

var instance *Server

func Initialize() {
	if !environment.Demo {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	if environment.OTelEnabled {
		r.Use(otelgin.Middleware(commons.ServiceName))
	}
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: environment.CORSAllowOrigins,
		AllowMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
		ExposeHeaders: []string{
			"Cache-Control", "Content-Language", "Content-Length", "Content-Type", "Expires", "Last-Modified", "Pragma",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	fiscusRouter := r.Group(environment.HTTPBasePath)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", environment.Port),
		Handler:           r,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       20 * time.Second,
		MaxHeaderBytes:    8 << 10, // 8 kilobytes
	}

	instance = &Server{
		router:       r,
		fiscusRouter: fiscusRouter,
		httpServer:   httpServer,
	}
}

func GetRouter() *gin.Engine {
	return instance.router
}

func GetFiscusRouter() *gin.RouterGroup {
	return instance.fiscusRouter
}

func Start() <-chan error {
	var serverErr = make(chan error, 1)
	go func() {
		slog.Debug("starting server")
		err := instance.httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Debug("server stopped unexpectedly")
			serverErr <- err
			return
		}
		slog.Debug("server stopped gracefully")
	}()
	return serverErr
}

func Shutdown(ctx context.Context) {
	if err := instance.httpServer.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "err", err)
	} else {
		slog.Info("server shutdown successful")
	}
}
