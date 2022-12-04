package repository

import (
	"genshinacademycore/config"
	"genshinacademycore/logger"
	models_db "genshinacademycore/models/db"

	"gorm.io/gorm"
)

const (
	PreloadCharacterName string = "Name"
	PreloadStatsProfit   string = "StatsProfit"
	PreloadElement       string = "Element"
	PreloadSlot          string = "Slot"
)

var (
	CharacterAllPreloads []string = []string{PreloadCharacterName, PreloadStatsProfit}
)

type FerretRepositoryInterface interface {
	GetAllCharactersStats() (*[]models_db.Character, error)
	// GetAllCharacters(preloads ...string) (*[]models.Character, error)
	// GetCharacter(id string, preloads ...string) (*[]models.Character, error)
	CreateTx() *gorm.DB
	RollbackTx() *gorm.DB
	CommitTx() error
}

type ferretRepository struct {
	DB *gorm.DB
}

func NewFerretRepository(dbConfig config.Database) FerretRepositoryInterface {
	return &ferretRepository{
		DB: dbConfig.ORM,
	}
}

func (repo *ferretRepository) GetAllCharactersStats() (*[]models_db.Character, error) {
	characters := &[]models_db.Character{}
	if err := repo.DBWithPreloads(CharacterAllPreloads).Find(characters).Error; err != nil {
		logger.Log.Error("Error GetAllCharacters")
		return nil, err
	}
	return characters, nil
}

// func (p *FerretRepository) GetAllCharacters(preloads ...string) (*[]models.Character, error) {
// 	character := &[]models.Character{}
// 	if err := p.DBWithPreloads(preloads).Find(character).Error; err != nil {
// 		logger.Log.Error("Error GetAllCharacters")
// 		return nil, err
// 	}
// 	return character, nil
// }

// func (p *FerretRepository) GetCharacter(id string, preloads ...string) (*[]models.Character, error) {
// 	character := &[]models.Character{}
// 	if err := p.DBWithPreloads(preloads).Where("id = ?", id).Find(character).Error; err != nil {
// 		logger.Log.Error("Error GetCharacter")
// 		return nil, err
// 	}
// 	return character, nil
// }

// DBWithPreloads - preload data.
func (p *ferretRepository) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := p.DB

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}

// CreateTx - create database transaction
func (p *ferretRepository) CreateTx() *gorm.DB {
	return p.DB.Begin()
}

// RollbackTx - rollback transaction
func (p *ferretRepository) RollbackTx() *gorm.DB {
	return p.DB.Rollback()
}

// CommitTx - commit transaction
func (p *ferretRepository) CommitTx() error {
	return p.DB.Commit().Error
}
