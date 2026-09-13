package handler

import (
	"context"
	"log"

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

func badLogOnly(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	ctx, span := tracer.Start(ctx, "badLogOnly")
	defer span.End()

	err := doWork(ctx)
	// ruleid: tqa-silent-error-not-recorded-go
	if err != nil {
		log.Printf("work failed: %v", err)
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

// goodDeferredRecord uses a named return value so the deferred closure
// can inspect the final error and record it on every return path,
// regardless of where in the function the error occurs.
func goodDeferredRecord(ctx context.Context) (err error) {
	tracer := otel.Tracer("svc")
	ctx, span := tracer.Start(ctx, "goodDeferredRecord")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	// ok: tqa-silent-error-not-recorded-go
	if err := doWork(ctx); err != nil {
		return err
	}
	return nil
}