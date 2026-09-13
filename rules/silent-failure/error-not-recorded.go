package handler

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func badNoRecord(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	ctx, span := tracer.Start(ctx, "badNoRecord")
	defer span.End()

	// ruleid: tqa-silent-error-not-recorded-go
	if err := doWork(ctx); err != nil {
		return err
	}
	return nil
}

func goodBoth(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	ctx, span := tracer.Start(ctx, "goodBoth")
	defer span.End()

	// ok: tqa-silent-error-not-recorded-go
	if err := doWork(ctx); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func goodNoSpan(ctx context.Context) error {
	// ok: tqa-silent-error-not-recorded-go
	if err := doWork(ctx); err != nil {
		return err
	}
	return nil
}