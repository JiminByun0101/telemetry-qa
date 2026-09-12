package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
)

var paymentAttempts = prometheus.NewCounterVec(
	prometheus.CounterOpts{Name: "payment_attempts_total"},
	[]string{"method", "status"},
)

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