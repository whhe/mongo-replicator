package redis

import (
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/whhe/mongo-replicator/token"
)

// manager store the latest resume token into redis.
//
// the token is stored as a string in extended json format. For details about extended json,
// see https://docs.mongodb.com/manual/reference/mongodb-extended-json/.
type manager struct {
	conn redis.Conn
	key  string
}

// NewTokenManager creates a redis token manager instance.
func NewTokenManager(uri string, key string) (token.Manager, error) {
	conn, err := redis.DialURL(uri)
	if err != nil {
		return nil, err
	}
	return &manager{conn, key}, nil
}

// Get gets the latest resume token.
func (r *manager) Get() (bson.Raw, error) {
	b, err := redis.Bytes(r.conn.Do("GET", r.key))
	if err != nil {
		return nil, err
	}
	var raw bson.Raw
	err = bson.UnmarshalExtJSON(b, true, &raw)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// Set saves the latest resume token.
func (r *manager) Set(b bson.Raw) error {
	_, err := r.conn.Do("SET", r.key, b.String())
	return err
}

// Close closes the redis connection.
func (r *manager) Close() {
	_ = r.conn.Close()
}
