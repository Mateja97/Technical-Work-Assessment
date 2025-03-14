package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	BalanceRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "balance_requests_total",
		Help: "Total number of balance requests",
	})

	BalanceRequestErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "balance_request_errors_total",
		Help: "Total number of failed balance requests",
	})

	BalanceRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "balance_request_duration_seconds",
		Help:    "Duration of balance requests in seconds",
		Buckets: prometheus.DefBuckets,
	})
)
