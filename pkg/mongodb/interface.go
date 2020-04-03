package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//OplogReader - Reads the records of the oplog collection
type OplogReader interface {
	Read(start time.Time) (*mongo.Cursor, error)
}

//Replicator - Starts the replication process
type Replicator interface {
	Start() error
}
