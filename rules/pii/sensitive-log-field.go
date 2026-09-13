package handler

import (
	"log/slog"

	"go.uber.org/zap"
)

func loginBad(email, password string, logger *zap.Logger) {
	// ruleid: tqa-pii-sensitive-log-field-go
	slog.Info("login attempt", slog.String("password", password))

	// ruleid: tqa-pii-sensitive-log-field-go
	logger.Info("auth", zap.Any("access_token", getToken()))
}

func loginGood(email, password string, logger *zap.Logger) {
	// ok: tqa-pii-sensitive-log-field-go
	slog.Info("login attempt", slog.String("password_hash", hash(password)))

	// ok: tqa-pii-sensitive-log-field-go
	logger.Info("auth", zap.String("token_type", "bearer"))

	// ok: tqa-pii-sensitive-log-field-go
	logger.Info("payment", zap.String("card_last4", last4(card)))

	// ok: tqa-pii-sensitive-log-field-go
	slog.Info("login ok", slog.String("account_id", userID))
}