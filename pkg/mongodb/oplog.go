package mongodb

import (
	"context"

	"github.com/involvestecnologia/mole/models"
	"github.com/involvestecnologia/mole/pkg/storage"
)

type oplogReplicator struct {
	oplogReader OplogReader
	storage     storage.Storage
}

//NewOplogReplication - Creates an instance of the replication service
func NewOplogReplication(oplogReader OplogReader, storage storage.Storage) Replicator {
	return &oplogReplicator{
		oplogReader: oplogReader,
		storage:     storage,
	}
}

//Start - Starts the replication process
func (o *oplogReplicator) Start() error {

	startTime, err := o.storage.StartTime()
	if err != nil {
		return err
	}

	cursor, err := o.oplogReader.Read(startTime)
	if err != nil {
		return err
	}

	for cursor.Next(context.TODO()) {
		oplog := &models.Oplog{}
		if err := cursor.Decode(oplog); err != nil {
			return err
		}
		if err := o.storage.Add(oplog); err != nil {
			return err
		}
	}

	return nil
}
