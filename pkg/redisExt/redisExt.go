package redisExt

import (
	"context"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/go-redsync/redsync/v4"
	pdkRedis "github.com/paper-indonesia/pdk/v2/redisExt"
	"github.com/redis/go-redis/v9"
)

const ServiceKey = "snap-core:"

type IRedisExt interface {
	Client() *redis.Client
	Close() error
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Incr(ctx context.Context, name string) *redis.IntCmd
	Limiter() *redis_rate.Limiter

	// Redsync
	NewMutex(name string, options ...redsync.Option) *redsync.Mutex
}

type redisExt struct {
	client  *redis.Client
	rs      *redsync.Redsync
	limiter *redis_rate.Limiter
}

func New(config pdkRedis.Config, opts ...pdkRedis.OptionFunc) (IRedisExt, error) {
	redis, err := pdkRedis.New(config, opts...)
	if err != nil {
		return nil, err
	}

	return &redisExt{redis.Client, redis.Redsync, redis.Limiter}, nil
}

func (r *redisExt) Client() *redis.Client {
	return r.client
}

func (r *redisExt) Limiter() *redis_rate.Limiter {
	return r.limiter
}

func (r *redisExt) Close() error {
	return r.client.Close()
}

func (r *redisExt) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	deletedKeys := make([]string, len(keys))
	for i, key := range keys {
		deletedKeys[i] = ServiceKey + key
	}
	return r.client.Del(ctx, deletedKeys...)
}

func (r *redisExt) Get(ctx context.Context, key string) *redis.StringCmd {
	if ctx.Err() != nil {
		ctx = context.Background()
	}
	return r.client.Get(ctx, ServiceKey+key)
}

func (r *redisExt) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if ctx.Err() != nil {
		ctx = context.Background()
	}
	return r.client.Set(ctx, ServiceKey+key, value, expiration)
}

func (r *redisExt) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.client.SetNX(ctx, ServiceKey+key, value, expiration)
}

func (r *redisExt) NewMutex(name string, options ...redsync.Option) *redsync.Mutex {
	return r.rs.NewMutex(name, options...)
}

func (r *redisExt) Ping(ctx context.Context) *redis.StatusCmd {
	return r.client.Ping(ctx)
}

func (r *redisExt) Incr(ctx context.Context, name string) *redis.IntCmd {
	return r.client.Incr(ctx, ServiceKey+name)
}
