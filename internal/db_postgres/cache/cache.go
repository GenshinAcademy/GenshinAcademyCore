package cache

import (
	db_models "ga/internal/db_postgres/models"
	"sync"
)

type Cache struct {
	mutex            *sync.Mutex
	characterStrings map[db_models.DBKey]*CharacterStrings
	newsStrings      map[db_models.DBKey]*NewsStrings
	tableStrings     map[db_models.DBKey]*TableStrings
	weaponStrings    map[db_models.DBKey]*WeaponStrings
}

func MakeCache(size uint64) *Cache {
	var cachePtr = new(Cache)

	cachePtr.characterStrings = make(map[db_models.DBKey]*CharacterStrings, size)
	cachePtr.newsStrings = make(map[db_models.DBKey]*NewsStrings, size)
	cachePtr.tableStrings = make(map[db_models.DBKey]*TableStrings, size)
	cachePtr.weaponStrings = make(map[db_models.DBKey]*WeaponStrings, size)
	cachePtr.mutex = new(sync.Mutex)

	return cachePtr
}

// Lock locks cache for simngle goroutine
func (cache Cache) Lock() {
	cache.mutex.Lock()
}

// Unlock unlock cache
func (cache Cache) Unlock() {
	cache.mutex.Unlock()
}
