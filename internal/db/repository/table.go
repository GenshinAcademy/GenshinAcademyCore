package repository

import (
	"context"
	"fmt"
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"

	"gorm.io/gorm"
)

type TableMapper interface {
	MapFromEntity(input *entity.Table, output *models.Table, language types.Language) error
	MapMultilingualFromEntity(input *entity.Table, output *models.TableMultilingual) error
	MapMultilingualFromModel(input *models.TableMultilingual, output *entity.Table) error
}

type TableRepository struct {
	db          *gorm.DB
	TableMapper TableMapper
}

func NewTableRepository(db *gorm.DB, TableMapper TableMapper) *TableRepository {
	return &TableRepository{
		db:          db,
		TableMapper: TableMapper,
	}
}

func (r *TableRepository) FindAllTables(ctx context.Context, language types.Language, offset int, limit int, sort string) ([]models.Table, error) {
	var tables []entity.Table

	query := r.db.WithContext(ctx).
		Model(&entity.Table{}).
		Offset(offset).
		Limit(limit).
		Order("created_at " + sort)

	if err := query.Find(&tables).Error; err != nil {
		return nil, fmt.Errorf("failed to find all tables: %w", err)
	}

	result := make([]models.Table, 0, len(tables))
	for _, table := range tables {
		var n models.Table
		if err := r.TableMapper.MapFromEntity(&table, &n, language); err != nil {
			return nil, fmt.Errorf("failed to map table: %w", err)
		}
		result = append(result, n)
	}

	return result, nil
}

func (r *TableRepository) CreateTable(ctx context.Context, table *models.TableMultilingual) error {
	dbTable := new(entity.Table)
	if err := r.TableMapper.MapMultilingualFromModel(table, dbTable); err != nil {
		return fmt.Errorf("failed to map table: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Create(dbTable).Error; err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (r *TableRepository) UpdateTable(ctx context.Context, id types.TableId, table *models.TableMultilingual) error {
	dbTable := new(entity.Table)
	if err := r.TableMapper.MapMultilingualFromModel(table, dbTable); err != nil {
		return fmt.Errorf("failed to map table: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Model(&entity.Table{
			Model: gorm.Model{
				ID: uint(id),
			},
		}).
		Updates(&entity.Table{
			Title:       dbTable.Title,
			Description: dbTable.Description,
			IconUrl:     dbTable.IconUrl,
			RedirectUrl: dbTable.RedirectUrl,
		}).Error; err != nil {
		return fmt.Errorf("failed to update table: %w", err)
	}

	return nil
}

func (r *TableRepository) DeleteTable(ctx context.Context, id types.TableId, force bool) error {
	dbTable := new(entity.Table)
	dbTable.ID = uint(id)

	tx := r.db
	if force {
		tx = tx.Unscoped()
	}

	if err := tx.Delete(&dbTable).Error; err != nil {
		return fmt.Errorf("failed to delete table: %w", err)
	}

	return nil
}
