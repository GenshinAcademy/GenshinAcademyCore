package genshin

import (
	"ga/internal/db_postgres"
	genshin_repositories "ga/internal/db_postgres/repositories/genshin"

	core "ga/pkg/genshin_core"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"fmt"
)

var (
	ConnectionFormat string = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	database         postgresDatabase
)

type postgresDatabaseConnection struct {
	IsFree        bool
	ORMConnection *gorm.DB
}

type postgresDatabase struct {
	Configuration PostgresDatabaseConfiguration
	Connections   []postgresDatabaseConnection
	Logger        gormLogger.Interface
}

// PostgresDatabaseConfiguration represents configuration for Postgres database.
type PostgresDatabaseConfiguration struct {
	Host         string
	UserName     string
	UserPassword string
	Name         string
	Port         uint16
	ServerPort   uint16
}

// GetConnectionString returns connection string from configuration.
func (dbConfig PostgresDatabaseConfiguration) GetConnectionString() string {
	return fmt.Sprintf(ConnectionFormat,
		dbConfig.Host,
		dbConfig.UserName,
		dbConfig.UserPassword,
		dbConfig.Name,
		dbConfig.Port)
}

// createPostgresProvider creates a provider to operate with Genshin database.
//
// @param language - Language to operate with.
func createPostgresProvider(language languages.Language) repositories.RepositoryProvider {
	return genshin_repositories.CreatePostgresGenshinCoreProvider(
		database.Connections[0].ORMConnection,
		language,
		db_postgres.GetCache())
}

// ConfigurePostgresDB applies postgres repositories to GenshinCore,
func ConfigurePostgresDB(config *core.GenshinCoreConfiguration) {
	config.ProviderFunc = core.GetProviderFunc(createPostgresProvider)
}

// newConnection creates new gorm connection and adds to connections pool,
func newConnection() postgresDatabaseConnection {
	orm, err := gorm.Open(postgres.Open(database.Configuration.GetConnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var connection = postgresDatabaseConnection{
		IsFree:        true,
		ORMConnection: orm,
	}
	database.Connections = append(database.Connections, connection)

	return connection
}

// InitializePostgresDatabase initializes database.
func InitializePostgresDatabase(config PostgresDatabaseConfiguration) {
	database = postgresDatabase{
		Configuration: config,
		Connections:   make([]postgresDatabaseConnection, 0),
	}
	newConnection()
	db_postgres.MigrateDatabase(database.Connections[0].ORMConnection)
	db_postgres.InitializeCache()
}

// CleanupConnections closes all active connections.
// Should be called with defer in main thread, either should be executed on application close.
func CleanupConnections() {
	for _, conn := range database.Connections {
		connection, _ := conn.ORMConnection.DB()
		connection.Close()
	}
}
