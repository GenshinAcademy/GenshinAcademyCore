package repository

import (
	"genshinacademycore/config"
	"genshinacademycore/logger"
	"genshinacademycore/models"

	"gorm.io/gorm"
)

type FerretRepositoryInterface interface {
	GetAllCharactersStats(preloads ...string) (*[]models.CharacterArtifactStatsProfit, error)
	GetAllCharacters(preloads ...string) (*[]models.Character, error)
	GetCharacter(id string, preloads ...string) (*[]models.Character, error)
	CreateTx() *gorm.DB
	RollbackTx() *gorm.DB
	CommitTx() error
}

type FerretRepository struct {
	DB *gorm.DB
}

func NewFerretRepository(dbConfig config.Database) FerretRepositoryInterface {
	return &FerretRepository{
		DB: dbConfig.ORM,
	}
}

func (p *FerretRepository) GetAllCharactersStats(preloads ...string) (*[]models.CharacterArtifactStatsProfit, error) {
	character := &[]models.CharacterArtifactStatsProfit{}
	if err := p.DBWithPreloads(preloads).Find(character).Error; err != nil {
		logger.Log.Error("Error GetAllCharacters")
		return nil, err
	}
	return character, nil
}

func (p *FerretRepository) GetAllCharacters(preloads ...string) (*[]models.Character, error) {
	character := &[]models.Character{}
	if err := p.DBWithPreloads(preloads).Find(character).Error; err != nil {
		logger.Log.Error("Error GetAllCharacters")
		return nil, err
	}
	return character, nil
}

func (p *FerretRepository) GetCharacter(id string, preloads ...string) (*[]models.Character, error) {
	character := &[]models.Character{}
	if err := p.DBWithPreloads(preloads).Where("id = ?", id).Find(character).Error; err != nil {
		logger.Log.Error("Error GetCharacter")
		return nil, err
	}
	return character, nil
}

// DBWithPreloads - preload data.
func (p *FerretRepository) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := p.DB

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}

// CreateTx - create database transaction
func (p *FerretRepository) CreateTx() *gorm.DB {
	return p.DB.Begin()
}

// RollbackTx - rollback transaction
func (p *FerretRepository) RollbackTx() *gorm.DB {
	return p.DB.Rollback()
}

// CommitTx - commit transaction
func (p *FerretRepository) CommitTx() error {
	return p.DB.Commit().Error
}
