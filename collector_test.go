package replicator_test

import (
	"context"
	"log"

	"github.com/whhe/mongo-replicator"
)

var collector = &replicator.Collector{}

func doCollect(*replicator.Collector)         {}
func newCollectorFunc() *replicator.Collector { return collector }

func ExampleNewCollector() {
	// set source mongodb and watch scope
	uri := "mongodb://username@password@host:port"
	db := []string{"db"}
	coll := []string{"coll_1", "coll_2"}

	collector := replicator.NewCollector(uri, db, coll)

	doCollect(collector)
}

func ExampleCollector_Collect() {
	collector := newCollectorFunc()

	stream, err := collector.Collect()
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close(context.TODO())

	// ... do some work ...
}
