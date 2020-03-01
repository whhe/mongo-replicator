package main

import (
	"context"
	"log"

	"github.com/whhe/mongo-replicator/collector"
	"github.com/whhe/mongo-replicator/model"
	"github.com/whhe/mongo-replicator/operator/mongo"
	"github.com/whhe/mongo-replicator/replicator"
	"github.com/whhe/mongo-replicator/token"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// conf
	sourceUri := "mongodb://xxx"
	targetUri := "mongodb://yyy"
	db := []string{"db"}
	coll := []string{"collA", "collB"}
	redisUri := "redis://xxx"
	redisKey := "resumeToken"

	// setup resume token manager
	tokenManager, err := token.NewRedisManager(redisUri, redisKey)
	if err != nil {
		log.Fatal(err)
	}
	defer tokenManager.Close()

	// resume after existing token
	opt := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	if resumeToken, err := tokenManager.Get(); err == nil {
		opt.SetResumeAfter(resumeToken)
	}

	// set collector and fetch change streams
	stream, err := collector.New(sourceUri, db, coll).Collect(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close(context.TODO())

	// setup replicator
	op, err := mongo.NewOperator(targetUri)
	if err != nil {
		log.Fatal(err)
	}
	repl := replicator.New(op)

	// iterate the change stream and sync change event until the change stream is closed
	// by the server or there is an error getting the next event.
	for {
		if stream.TryNext(context.TODO()) {
			var event model.ChangeEvent
			if err := stream.Decode(&event); err != nil {
				log.Fatal(err)
			}
			log.Printf("change event: %+v", event)

			// replicate change event
			err = repl.Replicate(event)
			if err != nil {
				log.Fatal(err)
			}

			// store resume token
			err = tokenManager.Set(stream.ResumeToken())
			if err != nil {
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
