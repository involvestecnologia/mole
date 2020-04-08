package collectors

import "github.com/prometheus/client_golang/prometheus"

type OplogCollector struct {
	oplog_reading_counter prometheus.Counter
}

func NewOplogCollector() OplogCollector {
	collector := OplogCollector{
		oplog_reading_counter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_reading_counter",
				Help: "Shows the amount of Oplog read from MongoDB",
			},
		),
	}

	prometheus.MustRegister([]prometheus.Collector{collector.oplog_reading_counter}...)
	return collector
}

func (o *OplogCollector) IncreasesReadingMetrics() {
	o.oplog_reading_counter.Inc()
}
