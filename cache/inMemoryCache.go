package cache

type ICashe interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

type Cache struct {
	store map[string]interface{}
}

func New() Cache {
	return Cache{
		store: make(map[string]interface{}),
	}
}

func (cas *Cache) Get(key string) interface{} {
	return cas.store[key]
}

func (cas *Cache) Set(key string, value interface{}) {
	cas.store[key] = value
}

func (cas *Cache) Delete(key string) {
	delete(cas.store, key)
}
