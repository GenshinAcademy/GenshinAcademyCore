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
	connection.AutoMigrate(&db_models.Language{})
	connection.AutoMigrate(&db_models.String{})
	connection.AutoMigrate(&db_models.StringValue{})
	connection.AutoMigrate(&db_models.Character{})
	connection.AutoMigrate(&db_models.CharacterIcon{})
	connection.AutoMigrate(&db_models.ArtifactProfit{})
    connection.AutoMigrate(&db_models.News{})
	connection.AutoMigrate(&db_models.Table{})
}
