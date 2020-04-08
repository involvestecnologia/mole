package collectors

import "github.com/prometheus/client_golang/prometheus"

type StorageCollector struct {
	oplog_storage_counter prometheus.Counter
}

func NewStorageCollector() StorageCollector {

	return StorageCollector{
		oplog_storage_counter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_storage_counter",
				Help: "Shows the amount of Oplog stored in Elasticsearch",
			},
		),
	}
}

func (s *StorageCollector) IncreasesStorageMetrics(records int) {
	s.oplog_storage_counter.Add(float64(records))
}
