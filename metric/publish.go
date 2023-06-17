package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PublishedCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "gardener_mqtt_published_count",
		Help: "The total number of published messages",
	}, []string{"topic"})
	ReceivedPublishCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "gardener_mqtt_received_publish_count",
		Help: "The total number of received publish messages",
	}, []string{"topic"})
)
