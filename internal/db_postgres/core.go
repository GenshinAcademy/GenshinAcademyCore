package db_postgres

import (
	"ga/internal/db_postgres/cache"
	db_models "ga/internal/db_postgres/models"

	"gorm.io/gorm"
)

var (
	globalCache *cache.Cache
)

// InitializeCache initializes the global cache if it hasn t been initialized.
func InitializeCache() {
	if globalCache == nil {
		globalCache = cache.MakeCache(128)
	}
}

// GetCache returns the global cache.
func GetCache() *cache.Cache {
	return globalCache
}

// MigrateDatabase сreates database structure
func MigrateDatabase(connection *gorm.DB) {
	connection.AutoMigrate(&db_models.DbLanguage{})
	connection.AutoMigrate(&db_models.DbString{})
	connection.AutoMigrate(&db_models.DbStringvalue{})
	connection.AutoMigrate(&db_models.DbCharacter{})
	connection.AutoMigrate(&db_models.DbCharacterIcon{})
	connection.AutoMigrate(&db_models.DbArtifactProfit{})
}
