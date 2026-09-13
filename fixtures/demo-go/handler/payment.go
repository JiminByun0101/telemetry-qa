package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
)

var paymentAttempts = prometheus.NewCounterVec(
	prometheus.CounterOpts{Name: "payment_attempts_total"},
	[]string{"method", "status"},
)

type PaymentRequest struct {
	OrderID    string
	Amount     int64
	CardNumber string
}

// NOTE: C002 - timestamp used directly as a label value, new series every call
func recordPaymentBad() {
	paymentAttempts.WithLabelValues("card", time.Now().Format(time.RFC3339)).Inc()
}

// NOTE: C002 - same issue with a random UUID
func recordPaymentBad2() {
	paymentAttempts.WithLabelValues(uuid.New().String(), "ok").Inc()
}

// NOTE: clean - bounded values only
func recordPaymentGood(status string) {
	paymentAttempts.WithLabelValues("card", status).Inc()
}

// NOTE: C003 - raw request path used as attribute instead of route pattern
func tagRequestBad(r *http.Request) {
	attribute.String("http.target", r.URL.Path)
}

// NOTE: clean - normalized route pattern, not the raw path
func tagRequestGood(routePattern string) {
	attribute.String("http.route", routePattern)
}

// NOTE: P001 - card number logged directly via structured field
func logCardBad(cardNumber string) {
	slog.Info("charge attempt", slog.String("card_number", cardNumber))
}

// NOTE: P002 - request struct dumped whole; adding a field later
// silently starts leaking it through this exact line
func logRequestBad(req PaymentRequest) {
	log.Printf("processing payment request %+v", req)
}

// NOTE: clean - only non-sensitive fields selected explicitly
func logRequestGood(req PaymentRequest) {
	slog.Info("processing payment", slog.String("order_id", req.OrderID))
}