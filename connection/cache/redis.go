package cache

import (
	"context"
	"math"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/oculius/oculi/v2/common/encoding/json"
	"github.com/pkg/errors"
)

type (
	redisImpl struct {
		rdc    *redis.Client
		parser json.JSON
	}
)

func NewRedis(parser json.JSON, opts Options) (Cache, error) {
	redisOpts := redis.Options(opts)
	rdc := redis.NewClient(&redisOpts)
	result := &redisImpl{rdc, parser}
	if err := rdc.Ping(context.Background()).Err(); err != nil {
		return result, err
	}
	return result, nil
}

func (m *redisImpl) IsExists(ctx context.Context, key string) (bool, error) {
	ttl, err := m.rdc.TTL(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return ttl.Nanoseconds() >= -1, nil
}

func (m *redisImpl) TTL(ctx context.Context, key string) (int64, error) {
	ttl, err := m.rdc.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if ttl.Nanoseconds() <= 0 {
		return ttl.Nanoseconds(), nil
	}
	return int64(math.Round(ttl.Seconds())), nil
}

func (m *redisImpl) Ping(ctx context.Context) error {
	return m.rdc.Ping(ctx).Err()
}

func (m *redisImpl) Close() error {
	return m.rdc.Close()
}

func (m *redisImpl) GetOrSet(ctx context.Context, key string, data any, setter func() (any, error), ttl time.Duration) error {
	found, errGet := m.Get(ctx, key, data)
	if errGet != nil {
		return errGet
	}

	if !found {
		newData, errRunner := setter()
		if errRunner != nil {
			return errRunner
		}

		if errSet := m.Set(ctx, key, newData, ttl); errSet != nil {
			return errSet
		}

		if errTransfer := m.transfer(newData, data); errTransfer != nil {
			return errTransfer
		}
	}

	return nil
}

func (m *redisImpl) transfer(source any, dest any) error {
	buff, err := m.parser.Marshal(source)
	if err != nil {
		return err
	}

	if errParser := m.parser.Unmarshal(buff, dest); errParser != nil {
		return errParser
	}
	return nil
}

func (m *redisImpl) GetThenSet(ctx context.Context, key string, data any, setter func() (any, error), ttl time.Duration) (bool, error) {
	found, errGet := m.Get(ctx, key, data)
	if errGet != nil {
		return found, errGet
	}

	newData, errRunner := setter()
	if errRunner != nil {
		return found, errRunner
	}

	if errSet := m.Set(ctx, key, newData, ttl); errSet != nil {
		return found, errSet
	}

	return found, nil
}

const KeepTTL = redis.KeepTTL

func (m *redisImpl) Get(ctx context.Context, key string, data any) (bool, error) {
	buff, err := m.rdc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}

	if data != nil {
		if errParser := m.parser.Unmarshal([]byte(buff), data); errParser != nil {
			return false, errParser
		}
	}
	return true, nil
}

func (m *redisImpl) Set(ctx context.Context, key string, data any, ttl time.Duration) error {
	buff, errParser := m.parser.Marshal(data)
	if errParser != nil {
		return errParser
	}

	if err := m.rdc.Set(ctx, key, buff, ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (m *redisImpl) Client() *redis.Client {
	return m.rdc
}
