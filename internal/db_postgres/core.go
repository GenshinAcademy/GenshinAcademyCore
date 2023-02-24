package db_postgres

import (
	db_models "ga/internal/db_postgres/models"
    "ga/internal/db_postgres/cache"
	"gorm.io/gorm"
)

var (
    globalCache *cache.Cache
)

func InitializeCache() {
    if globalCache == nil {
        globalCache = cache.MakeCache(128)
    }
}

func GetCache() *cache.Cache {
    return globalCache
}

// MigrateDatabase Creates database structure
func MigrateDatabase(connection *gorm.DB) {
	connection.AutoMigrate(&db_models.DbLanguage{})
	connection.AutoMigrate(&db_models.DbString{})
	connection.AutoMigrate(&db_models.DbStringvalue{})
	connection.AutoMigrate(&db_models.DbCharacter{})
	connection.AutoMigrate(&db_models.DbCharacterIcon{})
    connection.AutoMigrate(&db_models.DbArtifactProfit{})
}
