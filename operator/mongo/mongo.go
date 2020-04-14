package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/whhe/mongo-replicator/model"
	"github.com/whhe/mongo-replicator/operator"
)

type mongoOperator struct {
	*mongo.Client
}

// NewOperator creates a mongoOperator instance.
func NewOperator(uri string) (operator.Operator, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &mongoOperator{client}, nil
}

func (m *mongoOperator) Insert(e model.ChangeEvent) error {
	_, err := m.Database(e.Namespace.Database).
		Collection(e.Namespace.Collection).
		InsertOne(context.Background(), e.FullDocument)
	return err
}

func (m *mongoOperator) Delete(e model.ChangeEvent) error {
	_, err := m.Database(e.Namespace.Database).
		Collection(e.Namespace.Collection).
		DeleteOne(context.Background(), e.DocumentKey)
	return err
}

func (m *mongoOperator) Replace(e model.ChangeEvent) error {
	_, err := m.Database(e.Namespace.Database).
		Collection(e.Namespace.Collection).
		ReplaceOne(context.Background(), e.DocumentKey, e.FullDocument)
	return err
}

func (m *mongoOperator) Update(e model.ChangeEvent) error {
	if e.FullDocument == nil {
		update := bson.M{}

		if len(e.UpdateDescription.UpdatedFields) != 0 {
			update["$set"] = e.UpdateDescription.UpdatedFields
		}

		if len(e.UpdateDescription.RemovedFields) != 0 {
			unset := bson.M{}
			for _, field := range e.UpdateDescription.RemovedFields {
				unset[field] = ""
			}
			update["$unset"] = unset
		}

		_, err := m.Database(e.Namespace.Database).
			Collection(e.Namespace.Collection).
			UpdateOne(context.Background(), e.DocumentKey, update)
		return err
	}
	return m.Replace(e)
}

func (m *mongoOperator) Drop(e model.ChangeEvent) error {
	return m.Database(e.Namespace.Database).
		Collection(e.Namespace.Collection).
		Drop(context.Background())
}

func (m *mongoOperator) Rename(e model.ChangeEvent) error {
	from := e.Namespace.Database + "." + e.Namespace.Collection
	to := e.To.Database + "." + e.To.Collection

	result := m.Database("admin").
		RunCommand(
			context.Background(),
			bson.D{{Key: "renameCollection", Value: from}, {Key: "to", Value: to}},
		)
	return result.Err()
}

func (m *mongoOperator) DropDatabase(e model.ChangeEvent) error {
	return m.Database(e.Namespace.Database).Drop(context.Background())
}

func (m *mongoOperator) Invalidate(e model.ChangeEvent) error {
	panic("watch invalidate")
}
