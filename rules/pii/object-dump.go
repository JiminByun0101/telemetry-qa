package handler

import (
	"fmt"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

type Request struct {
	UserID string
	Path   string
}

func logBad(r Request) {
	// ruleid: tqa-pii-object-dump-go
	log.Printf("handling request %+v", r)
}

func logBad2(r Request) {
	// ruleid: tqa-pii-object-dump-go
	msg := fmt.Sprintf("request: %+v", r)
	_ = msg
}

func logBad3(r Request, logger *zap.Logger) {
	// ruleid: tqa-pii-object-dump-go
	logger.Info("request", zap.Any("request", r))
}

func logGood(r Request) {
	// ok: tqa-pii-object-dump-go
	slog.Info("handling request", slog.String("path", r.Path))
}

func logGood2(status int) {
	// ok: tqa-pii-object-dump-go
	log.Printf("responded with status %d", status)
}