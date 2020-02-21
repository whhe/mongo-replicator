package token

import (
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/bson"
)

// Manager interface can be used to customize how to get and update latest resume token.
type Manager interface {
	Get() (bson.Raw, error)
	Set(bson.Raw) error
	Close()
}

// RedisManager store the latest resume token into redis.
//
// The token is stored as a string in extended json format. For details about extended json,
// see https://docs.mongodb.com/manual/reference/mongodb-extended-json/.
type RedisManager struct {
	conn redis.Conn
	key  string
}

func NewRedisManager(uri string, key string) (Manager, error) {
	conn, err := redis.DialURL(uri)
	if err != nil {
		return nil, err
	}
	return &RedisManager{conn, key}, nil
}

func (r *RedisManager) Get() (bson.Raw, error) {
	b, err := redis.Bytes(r.conn.Do("GET", r.key))
	if err != nil {
		return nil, err
	}
	var token bson.Raw
	err = bson.UnmarshalExtJSON(b, true, &token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *RedisManager) Set(b bson.Raw) error {
	_, err := r.conn.Do("SET", r.key, b.String())
	return err
}

func (r *RedisManager) Close() {
	_ = r.conn.Close()
}
