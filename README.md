# Mongo Replicator

[![Tag](https://img.shields.io/github/v/tag/whhe/mongo-replicator.svg)](https://github.com/whhe/mongo-replicator/releases)
[![GoDoc](https://godoc.org/github.com/whhe/mongo-replicator?status.svg)](https://godoc.org/github.com/whhe/mongo-replicator)
[![GitHub license](https://img.shields.io/github/license/whhe/mongo-replicator)](https://github.com/whhe/mongo-replicator/blob/master/LICENSE)
[![codebeat badge](https://codebeat.co/badges/c0fac453-c67c-490d-8b38-05c07a1a75e7)](https://codebeat.co/projects/github-com-whhe-mongo-replicator-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/whhe/mongo-replicator)](https://goreportcard.com/report/github.com/whhe/mongo-replicator)

Mongo Replicator is a real-time replication tool for MongoDB replica set. It use [Change Streams](https://docs.mongodb.com/manual/changeStreams/) to fetch real-time data changes, with which users can sync data to other containers in real time.

## Requirements

Change Stream require 3.6 or higher version of MongoDB, and is only available for replica set.

For watch scope, deployment or database level is available for 4.0 or higher version of MongoDB.

For ChangeStreamOptions, [startAtOperationTime](https://docs.mongodb.com/manual/changeStreams/#start-time) require MongoDB 4.0, and [startAfter](https://docs.mongodb.com/manual/changeStreams/#change-stream-start-after) require MongoDB 4.2.


## Concepts

Mongo Replicator consists of Collector, Replicator and Operator.

### Collector

Collector represents the data source to watch. 

You can set the watch scope when creating a Collector instance. According to the params, the data source can be a deployment (either a replica set or a sharded cluster), several databases or collections.

```go
NewCollector(uri string, databases []string, collections []string) *Collector 
```

Collector has a `Collect` method to perform the watch action.

```go
(c *Collector) Collect(opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
```

See [ChangeStreamOptions](https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#ChangeStreamOptions) to customize your change streams.

### Replicator

Replicator is the unified interface to perform the replication. It is initialized by an Operator and has only one method `Replicate`.

```go
func (r *Replicator) Replicate(e model.ChangeEvent) error 
```

### Operator

Operator interface defined the methods need to be implemented for the target container to deal with [change event](https://docs.mongodb.com/manual/reference/change-events/) documents.

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

You can implement your own Operator to customize the replication logic.

## Usage and Example

See [example](example_test.go) and [godoc](https://godoc.org/github.com/whhe/mongo-replicator) for reference.

## License

[Apache 2.0 License](LICENSE)
