package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fiscalize(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "hr fiscalization requested")
	// do work
	slog.InfoContext(ctx, "hr fiscalization finished")
	c.JSON(http.StatusOK, gin.H{"success": true})
}
