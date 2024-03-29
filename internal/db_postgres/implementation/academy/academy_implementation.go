package academy

import (
	"ga/internal/db_postgres"
	academy_repositories "ga/internal/db_postgres/repositories/academy"
	genshin_db_repositories "ga/internal/db_postgres/repositories/genshin"
	core "ga/pkg/genshin_core"
	"ga/pkg/genshin_core/models/languages"
	genshin_core_repositories "ga/pkg/genshin_core/repositories"

	academy_core "ga/internal/academy_core"
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"

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

type PostgresDatabaseConfiguration struct {
	Host         string
	UserName     string
	UserPassword string
	DatabaseName string
	Port         uint16
	ServerPort   uint16
}

// GetConnectionString returns connection string from configuration.
func (dbConfig PostgresDatabaseConfiguration) GetConnectionString() string {
	return fmt.Sprintf(ConnectionFormat,
		dbConfig.Host,
		dbConfig.UserName,
		dbConfig.UserPassword,
		dbConfig.DatabaseName,
		dbConfig.Port)
}

func createPostgresGenshinCoreProvider(language *languages.Language) genshin_core_repositories.RepositoryProvider {
	return genshin_db_repositories.CreatePostgresGenshinCoreProvider(
		database.Connections[0].ORMConnection,
		language,
		db_postgres.GetCache())
}

func createLanguageRepository() repositories.ILanguageRepository {
	return academy_repositories.CreatePostresLanguageRepository(database.Connections[0].ORMConnection)
}

// createPostgresProvider creates a provider to operate with Academy database.
//
// @param language - Language to operate with.
func createPostgresProvider(language *models.Language) repositories.IRepositoryProvider {
	return academy_repositories.CreateAcademyProvider(
		database.Connections[0].ORMConnection,
		language,
		db_postgres.GetCache())
}

// ConfigurePostgresDB applies postgres repositories to GenshinCore.
func ConfigurePostgresDB(config *academy_core.AcademyCoreConfiguration) {
	config.ProviderFunc = academy_core.GetAcademyProviderFunc(createPostgresProvider)
	config.LanguageRepoFunc = academy_core.GetLanguageRepositoryFunc(createLanguageRepository)

	//Genshin core related
	config.GenshinCoreConfiguration.ProviderFunc = core.GetProviderFunc(createPostgresGenshinCoreProvider)
}

// Creates new gorm connection and adds to connections pool.
func newConnection() (postgresDatabaseConnection, error) {
	orm, err := gorm.Open(postgres.Open(database.Configuration.GetConnectionString()), &gorm.Config{})
	if err != nil {
		return postgresDatabaseConnection{}, err
	}

	var connection = postgresDatabaseConnection{
		IsFree:        true,
		ORMConnection: orm,
	}
	database.Connections = append(database.Connections, connection)

	return connection, err
}

// InitializePostgresDatabase initializes database.
func InitializePostgresDatabase(config PostgresDatabaseConfiguration) error {
	database = postgresDatabase{
		Configuration: config,
		Connections:   make([]postgresDatabaseConnection, 0),
	}

	_, err := newConnection()
	if err != nil {
		return err
	}

	db_postgres.MigrateDatabase(database.Connections[0].ORMConnection)
	db_postgres.InitializeCache()
	return nil
}

// CleanupConnections closes all active connections.
// Should be called with defer in main thread, either should be executed on application close.
func CleanupConnections() {
	for _, conn := range database.Connections {
		connection, _ := conn.ORMConnection.DB()
		connection.Close()
	}
}
