package logging

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestHandleAddsTraceContextFromContext(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(traceContextHandler{inner: slog.NewJSONHandler(&buf, nil)})

	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    trace.TraceID{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
		SpanID:     trace.SpanID{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		TraceFlags: trace.FlagsSampled,
	})
	ctx := trace.ContextWithSpanContext(context.Background(), spanContext)

	logger.InfoContext(ctx, "inside a span")

	out := buf.String()
	if !strings.Contains(out, spanContext.TraceID().String()) {
		t.Errorf("log line missing trace_id %s: %s", spanContext.TraceID(), out)
	}
	if !strings.Contains(out, spanContext.SpanID().String()) {
		t.Errorf("log line missing span_id %s: %s", spanContext.SpanID(), out)
	}
}

func TestHandleWithoutSpanOmitsTraceContext(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(traceContextHandler{inner: slog.NewJSONHandler(&buf, nil)})

	logger.Info("outside a span")
	logger.InfoContext(context.Background(), "context without span")

	if strings.Contains(buf.String(), "trace_id") {
		t.Errorf("log line has trace_id without a span: %s", buf.String())
	}
}

func TestWithAttrsAndGroupPreserveTraceContext(t *testing.T) {
	var buf bytes.Buffer
	base := slog.New(traceContextHandler{inner: slog.NewJSONHandler(&buf, nil)})
	logger := base.With("component", "test").WithGroup("details")

	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    trace.TraceID{0xaa},
		SpanID:     trace.SpanID{0xbb},
		TraceFlags: trace.FlagsSampled,
	})
	ctx := trace.ContextWithSpanContext(context.Background(), spanContext)

	logger.InfoContext(ctx, "derived logger", "key", "value")

	if !strings.Contains(buf.String(), spanContext.TraceID().String()) {
		t.Errorf("derived logger lost trace enrichment: %s", buf.String())
	}
}
