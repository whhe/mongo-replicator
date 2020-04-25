package replicator_test

import (
	"log"

	"github.com/whhe/mongo-replicator"
	"github.com/whhe/mongo-replicator/operator/mongo"
)

func doReplicate(*replicator.Replicator) {}

func ExampleNewReplicator_mongo() {
	// create a mongo operator to be used by the replicator
	mo, err := mongo.NewOperator("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// create a replicator using mongo operator
	exampleReplicator := replicator.NewReplicator(mo)

	doReplicate(exampleReplicator)
}
