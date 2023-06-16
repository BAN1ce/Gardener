package main

import "github.com/BAN1ce/skyTree/metric"

func main() {

	metric.PublishCount.WithLabelValues("test").Inc()
}
