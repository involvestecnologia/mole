package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type OplogReader interface {
	Read(start time.Time) (*mongo.Cursor, error)
}

type Replicator interface {
	Start() error
}
