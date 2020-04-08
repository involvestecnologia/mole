package collectors

import "github.com/prometheus/client_golang/prometheus"

type OplogCollector struct {
	oplog_reading_counter prometheus.Counter
}

func NewOplogCollector() OplogCollector {
	return OplogCollector{
		oplog_reading_counter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_reading_counter",
				Help: "Shows the amount of Oplog read from MongoDB",
			},
		),
	}
}

func (o *OplogCollector) IncreasesReadingMetrics() {
	o.oplog_reading_counter.Inc()
}
