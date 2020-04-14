package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OpLog represents the document of operations log in "local.oplog.rs" collection.
// See https://docs.mongodb.com/manual/core/replica-set-oplog/
type OpLog struct {
	OpTime        primitive.Timestamp `bson:"ts" json:"ts"`
	Hash          int64               `bson:"h" json:"h"`
	Version       int                 `bson:"v" json:"v"`
	OperationType string              `bson:"op" json:"op"`
	Namespace     string              `bson:"ns" json:"ns"`
	WallClockTime time.Time           `bson:"wall" json:"wall"`
	Operation     bson.Raw            `bson:"o" json:"o"`
	Update        bson.Raw            `bson:"o2,omitempty" json:"o2,omitempty"`
}
