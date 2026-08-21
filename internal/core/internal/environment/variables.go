package environment

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/fiscusproject/fiscus/internal/core/commons"
)

var (
	LogLevel         = slog.LevelInfo
	LogFormat        = "json"
	Port             = 8888
	Demo             = false
	CORSAllowOrigins = []string{"*"}
	HTTPBasePath     = fmt.Sprintf("/%s", commons.ServiceName)
	OTelEnabled      = false
	OTelServiceName  = commons.ServiceName
)

func Load() {
	commons.LoadIntEnvVariable("FISCUS_PORT", &Port)

	var logLevel string
	commons.LoadStrEnvVariable("FISCUS_LOG_LEVEL", &logLevel)
	switch strings.ToLower(logLevel) {
	case "":
	case "debug":
		LogLevel = slog.LevelDebug
	case "info":
		LogLevel = slog.LevelInfo
	case "warn", "warning":
		LogLevel = slog.LevelWarn
	case "error":
		LogLevel = slog.LevelError
	default:
		slog.Warn(fmt.Sprintf("Log level '%s' not recognized, defaulting to 'info'", logLevel))
	}

	commons.LoadStrEnvVariable("FISCUS_LOG_FORMAT", &LogFormat)
	commons.LoadBoolEnvVariable("FISCUS_DEMO", &Demo)
	commons.LoadStrSliceEnvVariable("FISCUS_CORS_ALLOW_ORIGINS", &CORSAllowOrigins)
	commons.LoadStrEnvVariable("FISCUS_HTTP_BASE_PATH", &HTTPBasePath)
	commons.LoadBoolEnvVariable("FISCUS_OTEL_ENABLED", &OTelEnabled)
	commons.LoadStrEnvVariable("OTEL_SERVICE_NAME", &OTelServiceName)
}
