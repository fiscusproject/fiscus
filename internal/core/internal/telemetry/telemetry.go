package telemetry

import (
	"context"
	"log/slog"

	"github.com/fiscusproject/fiscus/internal/core/commons"
	"github.com/fiscusproject/fiscus/internal/core/internal/environment"
	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	tracerProvider *sdktrace.TracerProvider
	meterProvider  *sdkmetric.MeterProvider
)

func Initialize() {
	if !environment.OTelEnabled {
		return
	}

	ctx := context.Background()

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			attribute.String("service.name", environment.OTelServiceName),
			attribute.String("service.version", commons.Version),
		),
	)
	if err != nil {
		slog.Error("otel resource setup error", "error", err)
		res = resource.Default()
	}

	spanExporter, err := autoexport.NewSpanExporter(ctx)
	if err != nil {
		slog.Error("otel span exporter setup error", "error", err)
		return
	}

	tracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(spanExporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator())

	metricReader, err := autoexport.NewMetricReader(ctx)
	if err != nil {
		slog.Error("otel metric reader setup error", "error", err)
		return
	}

	meterProvider = sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(metricReader),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)
}

func Dispose(ctx context.Context) {
	if tracerProvider != nil {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			slog.Error("otel tracer provider shutdown error", "err", err)
		}
	}

	if meterProvider != nil {
		if err := meterProvider.Shutdown(ctx); err != nil {
			slog.Error("otel meter provider shutdown error", "err", err)
		}
	}
}
