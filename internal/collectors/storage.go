package collectors

import "github.com/prometheus/client_golang/prometheus"

//StorageCollector - Collector metrics
type StorageCollector struct {
	oplogStorageCounter prometheus.Counter
}

//NewStorageCollector - Build the metrics collector
func NewStorageCollector() StorageCollector {

	collector := StorageCollector{
		oplogStorageCounter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_storage_counter",
				Help: "Shows the amount of Oplog stored in Elasticsearch",
			},
		),
	}

	prometheus.MustRegister(collector.oplogStorageCounter)
	return collector
}

//IncreasesStorageMetrics - Increases the storage metric
func (s *StorageCollector) IncreasesStorageMetrics(records int) {
	s.oplogStorageCounter.Add(float64(records))
}
