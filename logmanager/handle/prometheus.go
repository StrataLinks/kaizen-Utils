package handle

import (
	"github.com/StrataLinks/kaizen-Utils/internal/logger"
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusHandle exposes log counts as Prometheus metrics.
type PrometheusHandle struct {
	counterVec *prometheus.CounterVec
}

// NewPrometheusHandle initializes a new Prometheus handler.
func NewPrometheusHandle() *PrometheusHandle {
	counterVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "log_messages",
			Help: "Count of log messages by level.",
		},
		[]string{"level"},
	)
	prometheus.MustRegister(counterVec)
	return &PrometheusHandle{
		counterVec: counterVec,
	}
}

// Log increments the appropriate Prometheus counter based on the log level.
func (p *PrometheusHandle) Log(level logger.LogLevel, message string) {
	p.counterVec.WithLabelValues(string(level)).Inc()
}
