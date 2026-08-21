package api

import (
	"github.com/fiscusproject/fiscus/internal/core/server"
)

func RegisterRoutes() {
	hrRouter := server.GetFiscusRouter().Group("/hr/v1")

	hrRouter.POST("/fiscalize", Fiscalize)
}
