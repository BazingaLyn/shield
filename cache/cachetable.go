package cache

import (
	"log"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	name  string
	items map[interface{}]*CacheItem

	cleanupTime     *time.Timer
	cleanupInterval time.Duration

	loggger *log.Logger

	loadData func(key interface{}, args ...interface{}) *CacheItem

	addedItem         []func(item *CacheItem)
	aboutToDeleteItem []func(item *CacheItem)
}

func (table *CacheTable) Count() int {
	table.RLock()
	defer table.RUnlock()
	return len(table.items)
}

func (table *CacheTable) Foreach(trans func(key interface{}, item *CacheItem)) {
	table.RLock()
	defer table.RUnlock()

	for k, v := range table.items {
		trans(k, v)
	}
}
