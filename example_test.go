package replicator_test

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/whhe/mongo-replicator"
	"github.com/whhe/mongo-replicator/model"
	"github.com/whhe/mongo-replicator/operator/mongo"
	"github.com/whhe/mongo-replicator/token/redis"
)

var changeEvent = model.ChangeEvent{}
var newResumeToken = bson.Raw{}

func ExampleNewCollector() {
	// set source mongodb and watch scope
	uri := "mongodb://username@password@host:port"
	db := []string{"db"}
	coll := []string{"coll_1", "coll_2"}

	// create a collector instance and start collecting change stream
	collector := replicator.NewCollector(uri, db, coll)
	stream, err := collector.Collect()
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close(context.TODO())

	// iterate the change stream and do some work with the change event document until
	// the change stream is closed by the server or there is an error getting the next event
	for {
		if stream.TryNext(context.TODO()) {
			var event model.ChangeEvent
			if err := stream.Decode(&event); err != nil {
				log.Fatal(err)
			}
			// ... do some work ...
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

func ExampleNewReplicator() {
	// create an mongo operator to be used by the replicator
	mo, err := mongo.NewOperator("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// create a replicator instance and replicate a change event
	repl := replicator.NewReplicator(mo)
	err = repl.Replicate(changeEvent)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNewTokenManager() {
	// set redis uri and resume token key
	uri := "redis://:password@host:port/db"
	key := "resumeToken"

	// create a token manager instance
	manager, err := redis.NewTokenManager(uri, key)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Close()

	// get the latest resume token
	resumeToken, err := manager.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("resume token:", resumeToken)

	// save the latest resume token
	err = manager.Set(newResumeToken)
	if err != nil {
		log.Fatal(err)
	}
}
