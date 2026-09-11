package handler

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func ListOrders(ctx context.Context) error {
	tracer := otel.Tracer("orders")
	ctx, span := tracer.Start(ctx, "ListOrders")
	defer span.End()

	if err := queryOrders(ctx); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	slog.Info("orders listed", slog.String("route", "/orders"))
	return nil
}