package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PublishCount = promauto.NewCounterVec(prometheus.CounterOpts{}, []string{"topic"})
)
