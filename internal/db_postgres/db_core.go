package db_postgres

import (
	db_models "ga/internal/db_postgres/models"

	"gorm.io/gorm"
)

func MigrateDatabase(connection *gorm.DB) {
	connection.AutoMigrate(&db_models.Db_String{})
	connection.AutoMigrate(&db_models.Db_Language{})
	connection.AutoMigrate(&db_models.Db_Character{})
	connection.AutoMigrate(&db_models.Db_CharacterIcon{})
}
