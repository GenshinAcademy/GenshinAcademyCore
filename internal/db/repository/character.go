package repository

import (
	"context"
	"fmt"
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"
	"gorm.io/gorm"
)

type CharacterMapper interface {
	MapFromEntity(input *entity.Character, output *models.Character, language types.Language) error
	MapMultilingualFromEntity(input *entity.Character, output *models.CharacterMultilingual) error
	MapMultilingualFromModel(input *models.CharacterMultilingual, output *entity.Character) error
	MapWeaselAppraiserCharacterFromEntity(input *entity.Character, output *models.WeaselAppraiserCharacter, language types.Language) error
}

type CharacterRepository struct {
	db              *gorm.DB
	characterMapper CharacterMapper
}

func NewCharacterRepository(
	db *gorm.DB,
	mapper CharacterMapper,
) *CharacterRepository {
	return &CharacterRepository{
		db:              db,
		characterMapper: mapper,
	}
}

func (r *CharacterRepository) FindAllByLanguageWithIcons(ctx context.Context, language types.Language, offset int, limit int) ([]models.Character, error) {
	var characters []entity.Character

	query := r.db.WithContext(ctx).
		Model(&entity.Character{}).
		Preload("Icons").
		Offset(offset).
		Limit(limit)

	if err := query.Find(&characters).Error; err != nil {
		return nil, fmt.Errorf("failed to find all characters: %w", err)
	}

	result := make([]models.Character, 0, len(characters))
	for _, character := range characters {
		var char models.Character
		if err := r.characterMapper.MapFromEntity(&character, &char, language); err != nil {
			return nil, fmt.Errorf("failed to map character: %w", err)
		}

		result = append(result, char)
	}

	return result, nil
}

func (r *CharacterRepository) FindAllByLanguageWithArtifactProfitsAndIcons(ctx context.Context, language types.Language, offset int, limit int) ([]models.WeaselAppraiserCharacter, error) {
	var characters []entity.Character

	query := r.db.WithContext(ctx).
		Model(&entity.Character{}).
		Preload("Icons").
		InnerJoins("ArtifactProfits").
		Offset(offset).
		Limit(limit)

	if err := query.Find(&characters).Error; err != nil {
		return nil, fmt.Errorf("failed to find all characters: %w", err)
	}

	result := make([]models.WeaselAppraiserCharacter, 0, len(characters))
	for _, character := range characters {
		var char models.WeaselAppraiserCharacter
		if err := r.characterMapper.MapWeaselAppraiserCharacterFromEntity(&character, &char, language); err != nil {
			return nil, fmt.Errorf("failed to map character: %w", err)
		}

		result = append(result, char)
	}

	return result, nil
}

func (r *CharacterRepository) FindAllMultilingual(ctx context.Context, offset int, limit int) ([]models.CharacterMultilingual, error) {
	var characters []entity.Character

	query := r.db.WithContext(ctx).
		Model(&entity.Character{}).
		Offset(offset).
		Limit(limit)

	if err := query.Find(&characters).Error; err != nil {
		return nil, fmt.Errorf("failed to find all characters: %w", err)
	}

	result := make([]models.CharacterMultilingual, 0, len(characters))
	for _, character := range characters {
		var char models.CharacterMultilingual
		if err := r.characterMapper.MapMultilingualFromEntity(&character, &char); err != nil {
			return nil, fmt.Errorf("failed to map character: %w", err)
		}

		result = append(result, char)
	}

	return result, nil
}

func (r *CharacterRepository) CreateCharacter(ctx context.Context, characterMultilingual *models.CharacterMultilingual) error {
	dbCharacter := new(entity.Character)
	if err := r.characterMapper.MapMultilingualFromModel(characterMultilingual, dbCharacter); err != nil {
		return fmt.Errorf("failed to map character: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Create(dbCharacter).Error; err != nil {
		return fmt.Errorf("failed to create character: %w", err)
	}

	return nil
}
