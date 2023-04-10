package main

import (
	academy "ga/internal/academy_core"
	db "ga/internal/db_postgres/implementation/academy"
	genshin "ga/pkg/genshin_core"

	"ga/internal/configuration"
	"ga/pkg/genshin_core/models/languages"

	"ga/internal/services/genshin/characters"
	"ga/internal/services/middlewares"
	"ga/internal/services/news"
	"ga/internal/services/tables"
	"ga/internal/services/weasel/appraiser"

	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	logger                 *zap.Logger
	weaselAppraiserService *appraiser.Service
	genshinService         *characters.Service
	newsService            *news.Service
	tablesService          *tables.Service
	env                    Config
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         uint16 `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	LogLevel       int8   `mapstructure:"LOG_LEVEL"`
	GinMode        string `mapstructure:"GIN_MODE"`
	SecretKey      string `mapstructure:"SECRET_KEY"`
	AssetsPath     string `mapstructure:"ASSETS_PATH"`
	AssetsHost     string `mapstructure:"ASSETS_HOST"`
}

func init() {

	cfg, err := configuration.New[Config]()
	if err != nil {
		panic(err)
	}

	env = cfg.ENV
	logger = configuration.GetLogger(env.LogLevel)

	var dbConfig db.PostgresDatabaseConfiguration = db.PostgresDatabaseConfiguration{
		Host:         env.DBHost,
		UserName:     env.DBUserName,
		UserPassword: env.DBUserPassword,
		DatabaseName: env.DBName,
		Port:         env.DBPort,
	}

	if err := db.InitializePostgresDatabase(dbConfig); err != nil {
		logger.Sugar().Panic(err)
	}

	//Initializing gacore ga_config and configure it for postgres db
	var ga_config academy.AcademyCoreConfiguration = academy.AcademyCoreConfiguration{
		GenshinCoreConfiguration: genshin.GenshinCoreConfiguration{
			DefaultLanguage: languages.English,
		},
		AssetsPath: env.AssetsHost,
	}
	db.ConfigurePostgresDB(&ga_config)

	// Create ga core
	gacore := academy.CreateAcademyCore(ga_config)

	// Create services
	weaselAppraiserService = appraiser.CreateService(gacore)
	genshinService = characters.CreateService(gacore)
	newsService = news.CreateService(gacore)
	tablesService = tables.CreateService(gacore)
}

// Web server here
func main() {
	defer db.CleanupConnections()
	defer logger.Sync()

	r := gin.Default()

	gin.SetMode(env.GinMode)

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
		characters.GET("/stats", weaselAppraiserService.GetAllCharactersWithProfits)
	}

	news := mainRoute.Group("/news")
	{
		news.GET("", newsService.GetAllNews)
		news.POST("/", middlewares.Authenticate(env.SecretKey), newsService.CreateNews)
		news.PATCH("/:id", middlewares.Authenticate(env.SecretKey), newsService.UpdateNews)
	}

	tables := mainRoute.Group("/tables")
	{
		tables.GET("/", tablesService.GetAllTables)
		tables.POST("/", middlewares.Authenticate(env.SecretKey), tablesService.CreateTable)
		tables.PATCH("/:id", middlewares.Authenticate(env.SecretKey), tablesService.UpdateTable)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	err := r.Run(":" + env.ServerPort)
	if err != nil {
		logger.Sugar().Panic(err)
	}
}
