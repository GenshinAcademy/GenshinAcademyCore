package cache

import db_models "ga/internal/db_postgres/models"

type Cache struct {
    characterStrings map[db_models.DBKey]*CharacterStrings
}

func MakeCache(size uint64) *Cache {
    var cachePtr = new(Cache)
    cachePtr.characterStrings = make(map[db_models.DBKey]*CharacterStrings, size)

    return cachePtr
}
