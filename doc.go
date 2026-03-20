// Package hystrixprometheus provides a Prometheus metrics collector for Hystrix (https://github.com/afex/hystrix-go)
// that can be used to export metrics to Prometheus.
// The collector is built on top of hystrix-go/hystrix/metric_collector interface.
//
//	 Usage
//
//		promC := hystrixprometheus.NewPrometheusCollector("hystrix", nil, prometheus.DefBuckets)
//		metricCollector.Registry.Register(promC.Collector)
//
// Deprecated: This package wraps the unmaintained github.com/afex/hystrix-go library.
// No maintained alternative with the same metricCollector.MetricCollector interface exists.
// Consider migrating to alternative circuit breaker libraries with native Prometheus support.
package hystrixprometheus
