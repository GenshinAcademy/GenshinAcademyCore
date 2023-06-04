package main

import (
	docs "ga/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	academy "ga/internal/academy_core"
	db "ga/internal/db_postgres/implementation/academy"
	genshin "ga/pkg/genshin_core"

	"ga/internal/configuration"
	"ga/pkg/genshin_core/models/languages"

	"ga/internal/services/assets"
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
	assetsService          *assets.Service
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
	assetsService = assets.CreateService(gacore, env.AssetsPath)
}

//	@BasePath					/api
//	@title						Genshin Academy Service API
//	@description				Genshin Academy API documentation
//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Token for endpoints
func main() {
	defer db.CleanupConnections()
	defer logger.Sync()

	r := gin.Default()

	gin.SetMode(env.GinMode)

	// TODO: Move all router related code to internal/router package
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Accept-Languages"},
		MaxAge:          12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	mainRoute := r.Group("/api")

	characters := mainRoute.Group("/characters")
	{
		characters.GET("/", middlewares.GetLimitOffset(), genshinService.GetAll)
		characters.POST("/", middlewares.Authenticate(env.SecretKey), genshinService.Create)
		characters.GET("/stats", middlewares.GetLimitOffset(), weaselAppraiserService.GetAll)
		characters.PATCH("/stats/:id", middlewares.Authenticate(env.SecretKey), weaselAppraiserService.UpdateStats)
	}

	news := mainRoute.Group("/news")
	{
		news.GET("/", middlewares.GetLimitOffset(), newsService.GetAll)
		news.POST("/", middlewares.Authenticate(env.SecretKey), newsService.Create)
		news.PATCH("/:id", middlewares.Authenticate(env.SecretKey), newsService.Update)
	}

	tables := mainRoute.Group("/tables")
	{
		tables.GET("/", middlewares.GetLimitOffset(), tablesService.GetAll)
		tables.POST("/", middlewares.Authenticate(env.SecretKey), tablesService.Create)
		tables.PATCH("/:id", middlewares.Authenticate(env.SecretKey), tablesService.Update)
	}

	assets := mainRoute.Group("/assets")
	{
		assets.POST("/*path", middlewares.Authenticate(env.SecretKey), assetsService.Upload)
		assets.DELETE("/*path", middlewares.Authenticate(env.SecretKey), assetsService.Delete)
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
