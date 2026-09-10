from prometheus_client import Counter, Histogram, Gauge

http_requests_bad = Counter(
    "http_requests_total",
    "Total HTTP requests",
    # ruleid: tqa-card-high-card-label-py
    ["method", "status", "user_id"],
)

latency_bad = Histogram(
    "http_duration_seconds",
    "Request latency",
    # ruleid: tqa-card-high-card-label-py
    labelnames=["endpoint", "session_id"],
)

http_requests_good = Counter(
    "http_requests_total",
    "Total HTTP requests",
    # ok: tqa-card-high-card-label-py
    ["method", "status", "route"],
)

cache_good = Gauge(
    "cache_entries",
    "Cache entries",
    # ok: tqa-card-high-card-label-py
    labelnames=["cache_name", "user_tier"],
)