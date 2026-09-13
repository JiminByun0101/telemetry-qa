package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
)

func initBad(ctx context.Context) {
	// ruleid: tqa-silent-otel-init-error-ignored-go
	exporter, _ := otlptracegrpc.New(ctx)

	// ruleid: tqa-silent-otel-init-error-ignored-go
	res, _ := resource.New(ctx)

	_ = exporter
	_ = res
}

func initGood(ctx context.Context) error {
	// ok: tqa-silent-otel-init-error-ignored-go
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return err
	}

	// ok: tqa-silent-otel-init-error-ignored-go
	res, err := resource.New(ctx)
	if err != nil {
		return err
	}

	_ = exporter
	_ = res
	return nil
}