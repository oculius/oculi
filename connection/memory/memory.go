package memory

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/oculius/oculi/v2/connection"
	"time"
)

type (
	Memory interface {
		connection.Connection
		Client() *redis.Client
		Get(ctx context.Context, key string, data any) (bool, error)
		Set(ctx context.Context, key string, data any, ttl time.Duration) error
		// GetOrSet is a method to get from key, then if not found, set it on memory server and also set it on data.
		GetOrSet(ctx context.Context, key string, data any, setter func() (any, error), ttl time.Duration) error
		// GetThenSet is a method to get from key then put it on data, then do the
		// setter and set data on memory to the result of setter function with specified TTL.
		GetThenSet(ctx context.Context, key string, data any, setter func() (any, error), ttl time.Duration) (bool, error)
	}

	Options redis.Options
)
