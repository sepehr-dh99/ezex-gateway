package port

import (
	"context"
	"time"
)

type RedisPort interface {
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetJSON(ctx context.Context, key string, val any, ttl time.Duration) error
	GetJSON(ctx context.Context, key string, out any) error

	Close() error
}
