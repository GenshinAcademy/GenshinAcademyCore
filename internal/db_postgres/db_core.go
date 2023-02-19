package db_postgres

import (
	db_models "ga/internal/db_postgres/models"

	"gorm.io/gorm"
)

// MigrateDatabase Creates database structure
func MigrateDatabase(connection *gorm.DB) {
	connection.AutoMigrate(&db_models.DbLanguage{})
	connection.AutoMigrate(&db_models.DbString{})
	connection.AutoMigrate(&db_models.DbStringvalue{})
	connection.AutoMigrate(&db_models.DbCharacter{})
	connection.AutoMigrate(&db_models.DbCharacterIcon{})
    connection.AutoMigrate(&db_models.DbArtifactProfit{})
}
