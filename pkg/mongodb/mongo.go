package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/involvestecnologia/mole/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	conn *mongo.Client
}

//New
func New(config models.Mongo) OplogReader {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		log.Fatal(err)
	}

	return &db{
		conn: conn,
	}
}

func (d *db) Read(start time.Time) (*mongo.Cursor, error) {
	opts := options.Find()
	opts.SetCursorType(options.TailableAwait)
	opts.SetSort(bson.M{"$natural": 1})

	filter := bson.M{
		"ts": bson.M{
			"$gt": primitive.Timestamp{T: uint32(start.Unix())},
		},
		"ns": bson.M{
			"$ne": "config.system.sessions",
		},
		"op": bson.M{
			"$in": []string{"i", "u", "d"},
		},
	}

	return d.conn.Database("local").Collection("oplog.rs").Find(context.TODO(), filter, opts)
}
