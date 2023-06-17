package main

import (
	"context"
	"github.com/BAN1ce/gardener/client"
	"github.com/BAN1ce/gardener/logger"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2526", nil)

	go func() {
		var (
			topics = []string{"topic1", "topic2", "topic3", "topic4", "topic5"}
		)
		for i := 0; i < 100; i++ {
			go func() {
				client := client.Client{}
				client.Start(context.TODO())
				time.Sleep(1 * time.Second)
				for _, topic := range topics {
					if err := client.Subscribe(topic); err != nil {
						logger.Logger.Error("subscribe topic error", "topic = ", topic, "error = ", err)
					} else {
						logger.Logger.Info("subscribe topic", "topic = ", topic)
					}
				}
				for _, topic := range topics {
					client.Publish(1000, topic, []byte("hello world"))
				}
			}()
		}
	}()

	select {}
}
