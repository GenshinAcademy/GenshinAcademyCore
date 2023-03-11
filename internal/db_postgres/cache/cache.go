package cache

import (
    db_models "ga/internal/db_postgres/models"
    "sync"
)

type Cache struct {
    mutex *sync.Mutex
    characterStrings map[db_models.DBKey]*CharacterStrings
}

func MakeCache(size uint64) *Cache {
    var cachePtr = new(Cache)
    cachePtr.characterStrings = make(map[db_models.DBKey]*CharacterStrings, size)
    cachePtr.mutex = new(sync.Mutex)

    return cachePtr
}

// Lock locks cache for simngle goroutine
func(cache Cache) Lock() {
    cache.mutex.Lock()
}

// Unlock unlock cache
func (cache Cache) Unlock() {
    cache.mutex.Unlock()
}
