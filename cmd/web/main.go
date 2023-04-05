package main

import (
	"ga/internal/academy_core"
	"ga/internal/configuration"
	academy_postgres "ga/internal/db_postgres/implementation/academy"
	"ga/internal/genshin"
	"ga/internal/middlewares"
	"ga/internal/tables"
	"time"

	"ga/internal/news"
	core "ga/pkg/genshin_core"
	"ga/pkg/genshin_core/models/languages"
	"net/http"

	"ga/internal/ferret"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	logger         *zap.Logger
	ferretService  *ferret.FerretService
	genshinService *genshin.GenshinService
	newsService    *news.NewsService
	tablesService  *tables.TablesService
)

func init() {
	err := configuration.Init()
	if err != nil {
		panic(err)
	}

	logger = configuration.GetLogger()

	var dbConfig academy_postgres.PostgresDatabaseConfiguration = academy_postgres.PostgresDatabaseConfiguration{
		Host:         configuration.ENV.DBHost,
		UserName:     configuration.ENV.DBUserName,
		UserPassword: configuration.ENV.DBUserPassword,
		DatabaseName: configuration.ENV.DBName,
		Port:         configuration.ENV.DBPort,
	}

	if err := academy_postgres.InitializePostgresDatabase(dbConfig); err != nil {
		panic(err)
	}

	//Initializing gacore config and configure it for postgres db
	var config academy_core.AcademyCoreConfiguration = academy_core.AcademyCoreConfiguration{
		GenshinCoreConfiguration: core.GenshinCoreConfiguration{
			DefaultLanguage: languages.English,
		},
	}
	academy_postgres.ConfigurePostgresDB(&config)

	// Create ga core
	gacore := academy_core.CreateAcademyCore(config)

	// Create ferret service
	ferretService = ferret.CreateService(gacore)
	genshinService = genshin.CreateService(gacore)
	newsService = news.CreateService(gacore)
	tablesService = tables.CreateService(gacore)
}

// Web server here
func main() {
	defer academy_postgres.CleanupConnections()
	r := gin.Default()

	gin.SetMode(configuration.ENV.GinMode)

	// TODO: Move all router related code to internal/router package
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Accept-Language"},
		MaxAge:          12 * time.Hour,
	}))

	mainRoute := r.Group("/api")

	characters := mainRoute.Group("/characters")
	{
		characters.GET("/", genshinService.GetAllCharacters)
		characters.GET("/stats", ferretService.GetAllCharactersWithProfits)
	}

	news := mainRoute.Group("/news")
	{
		news.GET("", newsService.GetAllNews)
		news.POST("/", middlewares.Authenticate(configuration.ENV.SecretKey), newsService.CreateNews)
		news.PATCH("/:id", middlewares.Authenticate(configuration.ENV.SecretKey), newsService.UpdateNews)
	}

	tables := mainRoute.Group("/tables")
	{
		tables.GET("/", tablesService.GetAllTables)
		tables.POST("/", middlewares.Authenticate(configuration.ENV.SecretKey), tablesService.CreateTable)
		tables.PATCH("/:id", middlewares.Authenticate(configuration.ENV.SecretKey), tablesService.UpdateTable)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	err := r.Run(":" + configuration.ENV.ServerPort)
	if err != nil {
		logger.Sugar().Panic(err)
	}
}
