package main

import (
	"context"
	"ga/internal/config"
	assets_controller "ga/internal/controller/assets"
	characters_controller "ga/internal/controller/genshin/characters"
	news_controller "ga/internal/controller/news"
	tables_controller "ga/internal/controller/tables"
	appraiser_controller "ga/internal/controller/weasel/appraiser"
	"ga/internal/db"
	"ga/internal/router"
	assets_service "ga/internal/service/assets"
	character_service "ga/internal/service/characters"
	news_service "ga/internal/service/news"
	table_service "ga/internal/service/tables"
	appraiser_service "ga/internal/service/weasel/appraiser"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	cfg *config.Config
	r   *router.Router
)

func configureLogger(cfg config.LoggerConfig) {
	log.SetLevel(log.Level(cfg.LogLevel))
	log.SetOutput(os.Stdout)
	log.SetReportCaller(cfg.ReportCaller)
	switch cfg.LogFormatter {
	case 1:
		log.SetFormatter(&log.TextFormatter{})
	case 2:
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func init() {
	if err := enrichConfig(); err != nil {
		log.Fatal(err)
	}

	configureLogger(cfg.LoggerConfig)

	var dbConfig = &db.PostgresConfig{
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		Username:     cfg.DBUserName,
		Password:     cfg.DBUserPassword,
		DatabaseName: cfg.DBName,
	}

	p, err := db.NewPostgresProvider(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Migrate(context.Background()); err != nil {
		log.Fatal(err)
	}

	charRepo := p.GetCharacterRepository()
	newsRepo := p.GetNewsRepository()
	tableRepo := p.GetTableRepository()
	assetService := assets_service.New(cfg.AssetsPath, cfg.AssetsHost)
	r = router.New(cfg.GinMode, cfg.SecretKey).
		WithCharactersController(characters_controller.New(character_service.New(assetService, charRepo))).
		WithWeaselAppraiserController(appraiser_controller.New(appraiser_service.New(assetService, charRepo, p.GetArtifactProfitsRepository()))).
		WithNewsController(news_controller.New(news_service.New(assetService, newsRepo))).
		WithTablesController(tables_controller.New(table_service.New(assetService, tableRepo))).
		WithAssetsController(assets_controller.New(assetService))
}

func enrichConfig() (err error) {
	cfg, err = config.New()
	log.Infof("Config: %+v", cfg)
	return
}

// @BasePath					/api
// @title						Genshin Academy Service API
// @description				Genshin Academy Service API
// @securitydefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Token for endpoints
func main() {
	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
