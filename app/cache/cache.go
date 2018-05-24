package cache

import "time"

type item struct {
	expires int64
	value string
}

type Service struct {
	items map[string]item
	ttl int64
}

func Create(TTL int64) *Service {
	instance := new(Service)
	instance.ttl = TTL
	instance.items = make(map[string]item)
	return instance
}

func (cache *Service) Set(key string, value string) {
	cache.items[key] = item{
		expires: time.Now().Unix() + cache.ttl,
		value: value,
	}
}

func (cache *Service) Get(key string) (string, bool) {
	cache.cleanup()
	obj, ok := cache.items[key]
	if ok {
		return obj.value, true
	} else {
		return "", false
	}
}

func (cache *Service) cleanup() {
	now := time.Now().Unix()
	for k, v:= range cache.items {
		if v.expires < now {
			delete(cache.items, k)
		}
	}
}