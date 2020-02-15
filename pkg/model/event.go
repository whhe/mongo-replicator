package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// See https://docs.mongodb.com/manual/reference/change-events/
type ChangeEvent struct {
	ID                bson.Raw            `bson:"_id" json:"_id"`
	OperationType     string              `bson:"operationType" json:"operationType"`
	FullDocument      bson.Raw            `bson:"fullDocument,omitempty" json:"fullDocument,omitempty"`
	Namespace         Namespace           `bson:"ns" json:"ns"`
	To                Namespace           `bson:"to,omitempty" json:"to,omitempty"`
	DocumentKey       bson.Raw            `bson:"documentKey,omitempty" json:"documentKey,omitempty"`
	UpdateDescription bson.Raw            `bson:"updateDescription,omitempty" json:"updateDescription,omitempty"`
	ClusterTime       primitive.Timestamp `bson:"clusterTime" json:"clusterTime"`
	TransactionNumber int64               `bson:"txnNumber,omitempty" json:"txnNumber,omitempty"`
	SessionIdentifier bson.Raw            `bson:"lsid,omitempty" json:"lsid,omitempty"`
}

type Namespace struct {
	DB         string `bson:"db,omitempty" json:"db,omitempty"`
	Collection string `bson:"coll,omitempty" json:"coll,omitempty"`
}
