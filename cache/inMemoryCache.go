package cache

import (
	"errors"
	"sync"
	"time"
)

type ICashe interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

type Cache struct {
	store map[string]interface{}
	mu    *sync.Mutex // guards
}

func New() Cache {
	return Cache{
		store: make(map[string]interface{}),
		mu:    new(sync.Mutex),
	}
}

func (cas *Cache) Get(key string) (interface{}, error) {
	cas.mu.Lock()
	defer cas.mu.Unlock()

	value, ok := cas.store[key]
	if !ok {
		return nil, errors.New("value is not found")
	}

	return value, nil
}

func (cas *Cache) Set(key string, value interface{}, timeout time.Duration) {
	cas.mu.Lock()
	defer cas.mu.Unlock()

	cas.store[key] = value

	go func() {
		time.Sleep(timeout)
		cas.Delete(key)
	}()
}

func (cas *Cache) Delete(key string) {
	cas.mu.Lock()
	defer cas.mu.Unlock()

	delete(cas.store, key)
}
