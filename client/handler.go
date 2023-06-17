package client

import (
	"github.com/BAN1ce/gardener/logger"
	"github.com/BAN1ce/gardener/metric"
	"github.com/eclipse/paho.golang/packets"
	"github.com/eclipse/paho.golang/paho"
)

type Handler struct {
}

func (h *Handler) RegisterHandler(s string, handler paho.MessageHandler) {
	return
}

func (h *Handler) UnregisterHandler(s string) {
	return
}

func (h *Handler) Route(publish *packets.Publish) {
	metric.ReceivedPublishCount.WithLabelValues(publish.Topic).Inc()
	logger.Logger.Info("receive publish message", "payload = ", string(publish.Payload), "topic = ", publish.Topic)
	return
}

func (h *Handler) SetDebugLogger(l paho.Logger) {

}
