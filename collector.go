package replicator

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collector is a used to fetch change stream from a client, database or collection.
type Collector struct {
	uri      string
	db       string
	coll     string
	pipeline mongo.Pipeline
}

// NewCollector creates a new Collector instance.
func NewCollector(uri string, databases []string, collections []string) *Collector {
	var (
		db   string
		coll string
	)

	if len(databases) == 1 {
		db = databases[0]
		if len(collections) == 1 {
			coll = collections[0]
		}
	}

	return &Collector{
		uri:      uri,
		db:       db,
		coll:     coll,
		pipeline: newPipeline(databases, collections),
	}
}

func newPipeline(databases []string, collections []string) mongo.Pipeline {
	match := make([]bson.E, 0)
	if len(databases) > 0 {
		match = append(match, bson.E{Key: "ns.db", Value: bson.M{"$in": databases}})
	}
	if len(collections) > 0 {
		match = append(match, bson.E{Key: "ns.coll", Value: bson.M{"$in": collections}})
	}
	if len(match) > 0 {
		return mongo.Pipeline{bson.D{{Key: "$match", Value: match}}}
	}
	return mongo.Pipeline{}
}

// Collect returns a change stream for all changes on the corresponding container.
//
// The opts parameter can be used to specify options for change stream creation. See
// https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#ChangeStreamOptions for details.
func (c *Collector) Collect(opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.uri))
	if err != nil {
		return nil, err
	}
	// watch a client
	if c.db == "" {
		return client.Watch(context.Background(), c.pipeline, opts...)
	}
	// watch a database
	if c.coll == "" {
		return client.Database(c.db).Watch(context.Background(), c.pipeline, opts...)
	}
	// watch a collection
	return client.Database(c.db).Collection(c.coll).Watch(context.Background(), c.pipeline, opts...)
}
