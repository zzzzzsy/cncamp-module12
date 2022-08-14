package middleware

import "github.com/penglongli/gin-metrics/ginmetrics"

func GinMonitor() *ginmetrics.Monitor {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	return m
}
