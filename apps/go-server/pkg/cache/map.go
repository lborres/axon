package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

// In Memory Caching Implementation using go maps

const (
	defaultMaxSize = 1000
)

type cacheRecord[T any] struct {
	value     T
	expiresAt time.Time
	createdAt time.Time
}

type CacheConfig struct {
	MaxSize int // 0 = unlimited
}

type InMemoryCache[T any] struct {
	mu      sync.RWMutex
	records map[string]*cacheRecord[T]
	maxSize int
}

// Compile-time interface check
var _ Cache[any] = (*InMemoryCache[any])(nil)

func NewInMemoryCache[T any](config *CacheConfig) *InMemoryCache[T] {
	cache := &InMemoryCache[T]{
		records: make(map[string]*cacheRecord[T]),
	}

	if config != nil {
		cache.maxSize = config.MaxSize
	} else {
		cache.maxSize = defaultMaxSize
	}

	// Start cleanup goroutine
	go cache.cleanup()

	return cache
}

func (r *cacheRecord[T]) isExpired() bool {
	if r.expiresAt.IsZero() {
		return false // No TTL set
	}
	return time.Now().After(r.expiresAt)
}

// WARN: Thread-unsafe - Must be called with lock held
func (c *InMemoryCache[T]) unsafeEvictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, record := range c.records {
		if oldestKey == "" || record.createdAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = record.createdAt
		}
	}

	if oldestKey != "" {
		delete(c.records, oldestKey)
	}
}

// WARN: Thread-unsafe - Must be called with lock held
func (c *InMemoryCache[T]) unsafeLen() int {
	return len(c.records)
}

func (c *InMemoryCache[T]) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, record := range c.records {
			if record.isExpired() {
				delete(c.records, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *InMemoryCache[T]) Get(ctx context.Context, key string) (T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var zero T

	record, exists := c.records[key]
	if !exists {
		return zero, ErrKeyNotFound
	}

	if record.isExpired() {
		return zero, ErrExpired
	}

	return record.value, nil
}

func (c *InMemoryCache[T]) Set(ctx context.Context, key string, value T, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.records[key]

	if !exists && c.maxSize > 0 && c.unsafeLen() >= c.maxSize {
		c.unsafeEvictOldest()
	}

	var expiresAt time.Time
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	c.records[key] = &cacheRecord[T]{
		value:     value,
		expiresAt: expiresAt,
		createdAt: time.Now(),
	}

	return nil
}

func (c *InMemoryCache[T]) Exists(ctx context.Context, key string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.records[key]
	if !exists {
		return false, nil
	}

	return !item.isExpired(), nil
}

func (c *InMemoryCache[T]) Delete(ctx context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.records, key)
	return nil
}

func (c *InMemoryCache[T]) Clear(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.records = make(map[string]*cacheRecord[T])
	return nil
}

// Thread-safe Length check
func (c *InMemoryCache[T]) Len(ctx context.Context) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.records)
}

func (c *InMemoryCache[T]) ApproxSizeBytes(ctx context.Context) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	total := 0
	for _, record := range c.records {
		b, _ := json.Marshal(record.value)
		total += len(b)
	}
	return total
}
