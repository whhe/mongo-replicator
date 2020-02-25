# Mongo Replicator

Mongo Replicator is a real-time replication tool for MongoDB replica set. It use [Change Streams](https://docs.mongodb.com/manual/changeStreams/) to fetch real-time data changes, with which users can sync data to other containers in real time.

## Requirements

Change Stream require 3.6 or higher version of MongoDB, and is only available for replica set.

For watch scope, deployment or database level is available for 4.0 or higher version of MongoDB.

For ChangeStreamOptions, [startAtOperationTime](https://docs.mongodb.com/manual/changeStreams/#start-time) require MongoDB 4.0, and [startAfter](https://docs.mongodb.com/manual/changeStreams/#change-stream-start-after) require MongoDB 4.2.


## Concepts

Mongo Replicator is built on structure of Collector, Replicator and Operator.

### Collector

Collector represent the data source to watch. 

You can set the watch scope when creating a Collector instance. According to the params, the data source can be a deployment (either a replica set or a sharded cluster), several databases or collections.

```go
New(uri string, databases []string, collections []string) *Collector 
```

The Collector instance has a `Collect` method to perform the watch action.

```go
(c *Collector) Collect(opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
```

See [ChangeStreamOptions](https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#ChangeStreamOptions) to customize your change streams.

### Replicator

Replicator is the unified interface to perform the replication. It is initialized by a Operator and has only one method `Replicate`.

```go
func (r *Replicator) Replicate(e model.ChangeEvent) error 
```

### Operator

Operator defined the functions need to be implemented.

```go
type Operator interface {
	Insert(model.ChangeEvent) error
	Delete(model.ChangeEvent) error
	Replace(model.ChangeEvent) error
	Update(model.ChangeEvent) error
	Drop(model.ChangeEvent) error
	Rename(model.ChangeEvent) error
	DropDatabase(model.ChangeEvent) error
	Invalidate(model.ChangeEvent) error
}
```

You can implement your own Operator and use it to customize the replication of changes.

## Getting Started

Include Mongo Replicator in your project:

```go
import "github.com/whhe/mongo-replicator"
```

See [example.go](example/example.go) for reference.

## License

Mongo Replicator is released under the Apache 2.0 license. See [LICENSE](LICENSE)
