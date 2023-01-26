package db

import (
	"fmt"
	"ga/internal/db_postgres"
	db_repositories "ga/internal/db_postgres/repositories"
	"ga/pkg/core"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	ConnectionFormat string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai"
	database         postgresDatabaseProvider
)

type postgresDatabaseConnection struct {
	IsFree        bool
	ORMConnection *gorm.DB
}

type postgresDatabaseProvider struct {
	Configuration PostgresDatabaseConfiguration
	Connections   []postgresDatabaseConnection
	Logger        gormLogger.Interface
}

// Configuration for postgres database
type PostgresDatabaseConfiguration struct {
	Host         string
	UserName     string
	UserPassword string
	Name         string
	Port         string
	ServerPort   string
}

// Gets connectionstring from configuration
func (dbConfig PostgresDatabaseConfiguration) GetConnectionString() string {
	return fmt.Sprintf(ConnectionFormat,
		dbConfig.Host,
		dbConfig.UserName,
		dbConfig.UserPassword,
		dbConfig.Name,
		dbConfig.Port)
}

func CreatePostgresProvider(language models.Language) repositories.IRepositoryProvider {
	return db_repositories.PostgresRepositoryProvider{
		GormConnection: database.Connections[0].ORMConnection,
		Language:       language,
	}
}

// Creates new postgres repository for working with languages
func CreatePostgresLanguageRepository() repositories.ILanguageRepository {
	return db_repositories.CreatePostresLanguageRepository(database.Connections[0].ORMConnection)
}

// Applies postgres repositories to GenshinCore
func ConfigurePostgresDB(config *core.GenshinCoreConfiguration) {
	config.ProviderFunc = core.GetProviderFunc(CreatePostgresProvider)
	config.LanguageRepoFunc = core.GetLanguageRepoFunc(CreatePostgresLanguageRepository)
}

// Creates new gorm connection and adds to connections pool
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

// Initializes database
func InitializePostgresDatabase(config PostgresDatabaseConfiguration) {
	database = postgresDatabaseProvider{
		Configuration: config,
		Connections:   make([]postgresDatabaseConnection, 0),
	}
	newConnection()
	db_postgres.MigrateDatabase(database.Connections[0].ORMConnection)
}

// Closes all active connections. Should be called with defer in main thread, either should be executed on application close
func CleanupConnections() {
	for _, conn := range database.Connections {
		connection, _ := conn.ORMConnection.DB()
		defer connection.Close()
	}
}
