package tracer

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/fx"
	"time"
)

const (
	service     = "product-api"
	environment = "development"
	id          = 1
)

// NewTracer returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func NewTracer() (*tracesdk.TracerProvider, error) {
	url := fmt.Sprintf("http://host.docker.internal:14268/api/traces")
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func StartTracer(l fx.Lifecycle, tracer *tracesdk.TracerProvider) {
	l.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register our TracerProvider as the global so any imported
				// instrumentation in the future will default to using it.
				otel.SetTracerProvider(tracer)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				// Do not make the application hang when it is shutdown.
				ctx, cancel := context.WithTimeout(ctx, time.Second*5)
				defer cancel()
				if err := tracer.Shutdown(ctx); err != nil {
					return err
				}
				return nil
			},
		},
	)
}
