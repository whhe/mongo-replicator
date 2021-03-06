package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChangeEvent represents the document of change stream.
// See https://docs.mongodb.com/manual/reference/change-events/
type ChangeEvent struct {
	ID                bson.Raw  `bson:"_id" json:"_id"`
	OperationType     string    `bson:"operationType" json:"operationType"`
	FullDocument      bson.Raw  `bson:"fullDocument,omitempty" json:"fullDocument,omitempty"`
	Namespace         Namespace `bson:"ns" json:"ns"`
	To                Namespace `bson:"to,omitempty" json:"to,omitempty"`
	DocumentKey       bson.Raw  `bson:"documentKey,omitempty" json:"documentKey,omitempty"`
	UpdateDescription struct {
		UpdatedFields bson.M   `bson:"updatedFields,omitempty" json:"updatedFields,omitempty"`
		RemovedFields []string `bson:"removedFields,omitempty" json:"removedFields,omitempty"`
	} `bson:"updateDescription,omitempty" json:"updateDescription,omitempty"`
	ClusterTime       primitive.Timestamp `bson:"clusterTime" json:"clusterTime"`
	TransactionNumber int64               `bson:"txnNumber,omitempty" json:"txnNumber,omitempty"`
	SessionIdentifier bson.Raw            `bson:"lsid,omitempty" json:"lsid,omitempty"`
}

// Namespace contains database name and collection name.
type Namespace struct {
	Database   string `bson:"db,omitempty" json:"db,omitempty"`
	Collection string `bson:"coll,omitempty" json:"coll,omitempty"`
}

func (ns Namespace) String() string {
	return fmt.Sprintf("%s.%s", ns.Database, ns.Collection)
}
