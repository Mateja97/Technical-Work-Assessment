package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	BalanceRequestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "balance_requests_total",
			Help: "Total number of balance requests",
		})

	BalanceRequestErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "balance_request_errors_total",
		Help: "Total number of failed balance requests",
	})

	BalanceRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "balance_request_duration_seconds",
		Help:    "Duration of balance requests in seconds",
		Buckets: prometheus.DefBuckets,
	})
)

func init() {
	prometheus.MustRegister(BalanceRequestsCounter, BalanceRequestDuration, BalanceRequestErrors)
}
