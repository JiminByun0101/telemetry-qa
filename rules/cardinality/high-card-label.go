package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var httpRequestsBad = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	// ruleid: tqa-card-high-card-label-go
	[]string{"method", "status", "user_id"},
)

var latencyBad = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{Name: "http_duration_seconds"},
	// ruleid: tqa-card-high-card-label-go
	[]string{"request_id"},
)

var httpRequestsGood = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	// ok: tqa-card-high-card-label-go
	[]string{"method", "status", "route"},
)

var queueDepthGood = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{Name: "queue_depth"},
	// ok: tqa-card-high-card-label-go
	[]string{"queue_name", "namespace"},
)

var cacheGood = prometheus.NewCounterVec(
	prometheus.CounterOpts{Name: "cache_ops_total"},
	// ok: tqa-card-high-card-label-go
	[]string{"cache_hit", "user_tier", "path_group"},
)