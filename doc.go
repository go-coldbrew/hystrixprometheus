// Package hystrixprometheus provides a Prometheus metrics collector for Hystrix (https://github.com/afex/hystrix-go)
// that can be used to export metrics to Prometheus.
// The collector is built on top of hystrix-go/hystrix/metric_collector interface.
//
//	 Usage
//
//		promC := hystrixprometheus.NewPrometheusCollector("hystrix", nil, prometheus.DefBuckets)
//		metricCollector.Registry.Register(promC.Collector)
package hystrixprometheus
