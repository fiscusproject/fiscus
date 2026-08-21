package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/fiscusproject/fiscus/internal/core/internal/environment"
	"go.opentelemetry.io/otel/trace"
)

func Initialize() {
	opts := &slog.HandlerOptions{Level: environment.LogLevel}

	format := strings.ToLower(environment.LogFormat)

	var handler slog.Handler
	if format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(traceContextHandler{inner: handler}))

	if format != "text" && format != "json" {
		slog.Warn(fmt.Sprintf("Log format '%s' not recognized, defaulting to 'json'", environment.LogFormat))
	}
}

type traceContextHandler struct {
	inner slog.Handler
}

func (h traceContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.inner.Enabled(ctx, level)
}

func (h traceContextHandler) Handle(ctx context.Context, record slog.Record) error {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		record.AddAttrs(
			slog.String("trace_id", spanContext.TraceID().String()),
			slog.String("span_id", spanContext.SpanID().String()),
		)
	}

	return h.inner.Handle(ctx, record)
}

func (h traceContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return traceContextHandler{inner: h.inner.WithAttrs(attrs)}
}

func (h traceContextHandler) WithGroup(name string) slog.Handler {
	return traceContextHandler{inner: h.inner.WithGroup(name)}
}
