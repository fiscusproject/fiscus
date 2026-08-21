package hr

import (
	"github.com/fiscusproject/fiscus/internal/hr/internal/api"
	"github.com/fiscusproject/fiscus/internal/hr/internal/environment"
)

func Initialize() {
	environment.Load()
	if !environment.HREnabled {
		return
	}

	api.RegisterRoutes()
}
