package mongodb

import (
	"context"

	"github.com/involvestecnologia/mole/internal/collectors"
	"github.com/involvestecnologia/mole/internal/storage"
	"github.com/involvestecnologia/mole/models"
	"github.com/sirupsen/logrus"
)

type oplogReplicator struct {
	oplogReader OplogReader
	storage     storage.Storage
	collector   collectors.OplogCollector
	log         *logrus.Logger
}

//NewOplogReplication - Creates an instance of the replication service
func NewOplogReplication(oplogReader OplogReader, storage storage.Storage, collector collectors.OplogCollector, log *logrus.Logger) Replicator {
	return &oplogReplicator{
		oplogReader: oplogReader,
		storage:     storage,
		collector:   collector,
		log:         log,
	}
}

//Start - Starts the replication process
func (o *oplogReplicator) Start() {

	startTime, err := o.storage.StartTime()
	if err != nil {
		o.log.Fatal(err)
	}

	cursor, err := o.oplogReader.Read(startTime)
	if err != nil {
		o.log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		oplog := &models.Oplog{}

		if err := cursor.Decode(oplog); err != nil {
			o.log.Fatal(err)
		}

		o.collector.IncreasesReadingMetrics()

		if err := o.storage.Add(oplog); err != nil {
			o.log.Fatal(err)
		}
	}
}
