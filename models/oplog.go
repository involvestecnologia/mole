package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Oplog struct {
	Timestamp   primitive.Timestamp    `bson:"ts" json:"ts"`
	HistoryID   int64                  `bson:"h" json:"h"`
	Operation   string                 `bson:"op" json:"op"`
	Namespace   string                 `bson:"ns" json:"ns"`
	Object      map[string]interface{} `bson:"o" json:"o"`
	QueryObject map[string]interface{} `bson:"o2" json:"o2"`
}

func (o *Oplog) GetOperationName() string {
	switch {
	case o.Operation == "i":
		return "insert"
	case o.Operation == "u":
		return "update"
	case o.Operation == "d":
		return "delete"
	default:
		return o.Operation
	}
}
