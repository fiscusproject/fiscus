package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Readyz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
