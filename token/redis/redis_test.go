package redis_test

import (
	"log"

	"github.com/whhe/mongo-replicator/token/redis"
)

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

	// ... get and set resume token ...
}
