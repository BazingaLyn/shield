package cache

import "sync"

var (
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

func Cache(table string) *CacheTable {
	mutex.RLock()
	cacheTable, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		cacheTable, ok = cache[table]
		if !ok {
			cacheTable = &CacheTable{
				name:  table,
				items: make(map[interface{}]*CacheItem),
			}
			cache[table] = cacheTable
		}
	}

	return cacheTable
}
