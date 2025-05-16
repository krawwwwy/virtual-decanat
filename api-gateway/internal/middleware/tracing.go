package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/zap"
	"context"
	"time"
)

// InitTracer инициализирует трассировку OpenTelemetry
func InitTracer(serviceName, collectorURL string, logger *zap.Logger) (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithEndpoint(collectorURL),
			otlptracegrpc.WithInsecure(),
		),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	logger.Info("Tracing initialized", zap.String("service", serviceName))
	return tp, nil
}

// TracingMiddleware добавляет трассировку запросов
func TracingMiddleware(serviceName string) gin.HandlerFunc {
	return otelgin.Middleware(serviceName)
}

// ShutdownTracer закрывает трассировщик
func ShutdownTracer(tp *sdktrace.TracerProvider) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		// Логирование ошибки закрытия трассировщика
	}
} 