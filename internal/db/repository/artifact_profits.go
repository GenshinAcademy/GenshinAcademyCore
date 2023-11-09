package repository

import (
	"context"
	"fmt"
	"ga/internal/db/entity"
	"ga/internal/types"
	"gorm.io/gorm"
)

type ArtifactProfitsRepository struct {
	db *gorm.DB
}

func NewArtifactProfitsRepository(
	db *gorm.DB,
) *ArtifactProfitsRepository {
	return &ArtifactProfitsRepository{
		db: db,
	}
}

func (r *ArtifactProfitsRepository) FindOneByCharacterId(ctx context.Context, id types.CharacterId) (types.CharacterArtifactProfits, error) {
	var profits = new(entity.ArtifactProfits)

	if err := r.db.WithContext(ctx).
		First(profits, id).Error; err != nil {
		return nil, fmt.Errorf("failed to find artifact profits: %w", err)
	}

	return profits.Profits, nil
}

func (r *ArtifactProfitsRepository) UpdateArtifactProfits(ctx context.Context, id types.CharacterId, artifactProfits types.CharacterArtifactProfits) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.ArtifactProfits{CharacterId: id}).
		Update("profits", artifactProfits).Error; err != nil {
		return fmt.Errorf("failed to update artifact profits: %w", err)
	}

	return nil
}
