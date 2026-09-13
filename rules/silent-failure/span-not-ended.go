package handler

import (
	"context"

	"go.opentelemetry.io/otel"
)

func badSimple(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	// ruleid: tqa-silent-span-not-ended-go
	ctx, span := tracer.Start(ctx, "badSimple")
	return doWork(ctx)
}

func goodDefer(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	// ok: tqa-silent-span-not-ended-go
	ctx, span := tracer.Start(ctx, "goodDefer")
	defer span.End()
	return doWork(ctx)
}

func goodExplicit(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	// ok: tqa-silent-span-not-ended-go
	ctx, span := tracer.Start(ctx, "goodExplicit")
	err := doWork(ctx)
	span.End()
	return err
}

func goodTwoSpans(ctx context.Context) error {
	tracer := otel.Tracer("svc")
	// ok: tqa-silent-span-not-ended-go
	ctx, outer := tracer.Start(ctx, "outer")
	defer outer.End()
	// ok: tqa-silent-span-not-ended-go
	ctx, inner := tracer.Start(ctx, "inner")
	defer inner.End()
	return doWork(ctx)
}