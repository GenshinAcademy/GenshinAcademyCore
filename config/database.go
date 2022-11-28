package config

import (
	"fmt"
	oldLog "log"
	"time"

	"genshinacademycore/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Database - .
type Database struct {
	ORM *gorm.DB
}

// InitDB - .
func InitDB(env *Config) Database {
	log := logger.Log.WithFields(logrus.Fields{
		"LstdFlags": oldLog.LstdFlags,
	})

	newLogger := gormLogger.New(
		log, // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.DBHost, env.DBUserName, env.DBUserPassword, env.DBName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Log.Fatal(err)
	}

	return Database{
		ORM: db,
	}
}
