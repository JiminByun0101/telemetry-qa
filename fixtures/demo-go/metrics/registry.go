package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// NOTE: user_id label - high cardinality, should be caught by C001
var HTTPRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "status", "user_id"},
)

// NOTE: clean - bounded labels, should NOT be flagged
var RequestLatency = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Request latency",
	},
	[]string{"route", "method"},
)

// newCounter wraps registration so services don't repeat boilerplate.
func newCounter(name string, labels []string) *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{Name: name},
		labels,
	)
}
 
// NOTE: same user_id issue as above, but registered through the wrapper -
// this one is expected to be MISSED by C001 (no interprocedural tracking)
var APICallsTotal = newCounter("api_calls_total", []string{"method", "user_id"})
 
