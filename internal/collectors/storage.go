package collectors

import "github.com/prometheus/client_golang/prometheus"

type StorageCollector struct {
	oplogStorageCounter prometheus.Counter
}

//NewStorageCollector
func NewStorageCollector() StorageCollector {

	collector := StorageCollector{
		oplogStorageCounter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_storage_counter",
				Help: "Shows the amount of Oplog stored in Elasticsearch",
			},
		),
	}

	prometheus.MustRegister([]prometheus.Collector{collector.oplogStorageCounter}...)
	return collector
}

func (s *StorageCollector) IncreasesStorageMetrics(records int) {
	s.oplogStorageCounter.Add(float64(records))
}
