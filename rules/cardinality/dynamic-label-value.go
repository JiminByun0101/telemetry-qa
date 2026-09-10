package metrics

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func recordBad(m *CounterVec) {
	// ruleid: tqa-card-timestamp-label-value-go
	m.WithLabelValues(time.Now().Format(time.RFC3339))

	// ruleid: tqa-card-timestamp-label-value-go
	m.WithLabelValues(uuid.New().String())

	// ruleid: tqa-card-timestamp-label-value-go
	m.WithLabelValues(strconv.FormatInt(time.Now().Unix(), 10))
}

func recordGood(m *CounterVec, route string, status int) {
	// ok: tqa-card-timestamp-label-value-go
	m.WithLabelValues(route, fmt.Sprint(status))
}
