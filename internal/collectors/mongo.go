package collectors

import "github.com/prometheus/client_golang/prometheus"

type OplogCollector struct {
	oplogReadingCounter prometheus.Counter
}

func NewOplogCollector() OplogCollector {
	collector := OplogCollector{
		oplogReadingCounter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_reading_counter",
				Help: "Shows the amount of Oplog read from MongoDB",
			},
		),
	}

	prometheus.MustRegister([]prometheus.Collector{collector.oplogReadingCounter}...)
	return collector
}

func (o *OplogCollector) IncreasesReadingMetrics() {
	o.oplogReadingCounter.Inc()
}
