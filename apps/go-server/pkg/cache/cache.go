package cache

import (
	"context"
	"errors"
	"time"
)

type Cache[T any] interface {
	Get(ctx context.Context, key string) (T, error)
	Set(ctx context.Context, key string, value T, ttl time.Duration) error

	Exists(ctx context.Context, key string) (bool, error)

	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error

	Len(ctx context.Context) int
	ApproxSizeBytes(ctx context.Context) int
}

var (
	ErrKeyNotFound = errors.New("cache: key not found")
	ErrExpired     = errors.New("cache: key expired")
	ErrCacheFull   = errors.New("cache: max capacity reached")
)
