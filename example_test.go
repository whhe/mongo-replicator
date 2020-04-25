package replicator_test

import (
	"context"
	"log"

	"github.com/whhe/mongo-replicator"
	"github.com/whhe/mongo-replicator/model"
)

var exampleReplicator = &replicator.Replicator{}

func newReplicatorFunc() *replicator.Replicator { return exampleReplicator }

func Example() {
	// create a collector
	collector := newCollectorFunc()

	// start collecting change streams
	stream, err := collector.Collect()
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close(context.TODO())

	// create a replicator
	repl := newReplicatorFunc()

	// iterate the change stream and sync change event until the change stream is
	// closed by the server or there is an error getting the next event document
	for {
		if stream.TryNext(context.TODO()) {
			var event model.ChangeEvent
			if err := stream.Decode(&event); err != nil {
				log.Fatal(err)
			}

			// replicate data using change event
			if err := repl.Replicate(event); err != nil {
				log.Fatal(err)
			}

			continue
		}

		if err := stream.Err(); err != nil {
			log.Fatal(err)
		}
		if stream.ID() == 0 {
			log.Print("the cursor has been closed or exhausted")
			break
		}
	}
}
