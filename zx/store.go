package zx

import (
	"errors"
	"sync"
	"time"
)

// Store is an interface for a key value store.
type Store interface {
	// Get retrieves a value from the store.
	Get(key string) ([]byte, error)
	// Set stores a value in the store.
	Set(key string, value []byte) error
	// Del deletes a value from the store.
	Del(key string) error
	// Exists checks if a key exists in the store.
	Exists(key string) bool
	// SetEx stores a value in the store with an expiry.
	SetEx(key string, value []byte, expiry time.Duration) error
	// Keys returns a list of keys in the store.
	Keys() []string
	// Close closes the store.
	Close() error
}

// ZStore is an in-memory key value store.
type ZStore struct {
	data       map[string]storeEntry
	mu         sync.RWMutex
	gcInterval time.Duration
	done       chan struct{}
}

// storeEntry is a key value pair with an optional expiry.
type storeEntry struct {
	data []byte
	exp  *time.Time
}

// NewZStore creates a new ZStore.
func NewZStore(gcInterval ...time.Duration) *ZStore {
	var duration time.Duration
	if len(gcInterval) > 0 {
		duration = gcInterval[0]
	} else {
		duration = time.Second * 15
	}

	z := &ZStore{
		data:       make(map[string]storeEntry),
		gcInterval: duration,
		mu:         sync.RWMutex{},
		done:       make(chan struct{}),
	}
	go z.gc()
	return z
}

func (z *ZStore) Close() error {
	z.done <- struct{}{}
	return nil
}

func (z *ZStore) Get(key string) ([]byte, error) {
	z.mu.RLock()
	defer z.mu.RUnlock()

	if v, ok := z.data[key]; ok {
		return v.data, nil
	}
	return nil, errors.New("key not found")
}

func (z *ZStore) Set(key string, value []byte) error {
	z.mu.Lock()
	defer z.mu.Unlock()

	z.data[key] = storeEntry{data: value}
	return nil
}

func (z *ZStore) Del(key string) error {
	z.mu.Lock()
	defer z.mu.Unlock()

	delete(z.data, key)
	return nil
}

func (z *ZStore) Exists(key string) bool {
	z.mu.RLock()
	defer z.mu.RUnlock()

	_, ok := z.data[key]
	return ok
}

func (z *ZStore) SetEx(key string, value []byte, expiry time.Duration) error {
	z.mu.Lock()
	defer z.mu.Unlock()

	exp := time.Now().Add(expiry)
	z.data[key] = storeEntry{data: value, exp: &exp}
	return nil
}

func (z *ZStore) Keys() []string {
	z.mu.RLock()
	defer z.mu.RUnlock()

	var keys []string
	for k := range z.data {
		keys = append(keys, k)
	}
	return keys
}

func (z *ZStore) gc() {
	ticker := time.NewTicker(z.gcInterval)
	defer ticker.Stop()

	for {
		select {
		case <-z.done:
			return
		case <-ticker.C:
			for k, v := range z.data {
				if v.exp != nil && v.exp.Before(time.Now()) {
					z.mu.Lock()
					delete(z.data, k)
					z.mu.Unlock()
				}
			}
		}
	}
}
