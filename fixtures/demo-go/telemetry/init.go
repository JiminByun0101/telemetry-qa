package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

// NOTE: S003 - exporter creation error discarded; if this fails the
// service still starts normally but emits no telemetry at all
func InitBad(ctx context.Context) {
	exporter, _ := otlptracegrpc.New(ctx)
	_ = exporter
}

// NOTE: clean - error is checked and propagated
func InitGood(ctx context.Context) error {
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return err
	}
	_ = exporter
	return nil
}

// NOTE: S001 - span started, never ended; this request's trace will
// never appear in the backend
func FetchAccountBad(ctx context.Context) error {
	tracer := otel.Tracer("accounts")
	ctx, span := tracer.Start(ctx, "FetchAccount")
	return lookupAccount(ctx)
}

// NOTE: clean - deferred end covers all return paths
func FetchAccountGood(ctx context.Context) error {
	tracer := otel.Tracer("accounts")
	ctx, span := tracer.Start(ctx, "FetchAccount")
	defer span.End()
	return lookupAccount(ctx)
}
