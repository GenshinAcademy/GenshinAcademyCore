package repository

import (
	"context"
	"fmt"
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"

	"gorm.io/gorm"
)

type NewsMapper interface {
	MapFromEntity(input *entity.News, output *models.News, language types.Language) error
	MapMultilingualFromEntity(input *entity.News, output *models.NewsMultilingual) error
	MapMultilingualFromModel(input *models.NewsMultilingual, output *entity.News) error
}

type NewsRepository struct {
	db         *gorm.DB
	newsMapper NewsMapper
}

func NewNewsRepository(db *gorm.DB, newsMapper NewsMapper) *NewsRepository {
	return &NewsRepository{
		db:         db,
		newsMapper: newsMapper,
	}
}

func (r *NewsRepository) FindAllNews(ctx context.Context, language types.Language, offset int, limit int, sort string) ([]models.News, error) {
	var news []entity.News

	query := r.db.WithContext(ctx).
		Model(&entity.News{}).
		Offset(offset).
		Limit(limit).
		Order("created_at " + sort)

	if err := query.Find(&news).Error; err != nil {
		return nil, fmt.Errorf("failed to find all news: %w", err)
	}

	result := make([]models.News, 0, len(news))
	for _, news := range news {
		var n models.News
		if err := r.newsMapper.MapFromEntity(&news, &n, language); err != nil {
			return nil, fmt.Errorf("failed to map news: %w", err)
		}
		result = append(result, n)
	}

	return result, nil
}

func (r *NewsRepository) CreateNews(ctx context.Context, news *models.NewsMultilingual) error {
	dbNews := new(entity.News)
	if err := r.newsMapper.MapMultilingualFromModel(news, dbNews); err != nil {
		return fmt.Errorf("failed to map news: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Create(dbNews).Error; err != nil {
		return fmt.Errorf("failed to create news: %w", err)
	}

	return nil
}

func (r *NewsRepository) UpdateNews(ctx context.Context, id types.NewsId, news *models.NewsMultilingual) error {
	dbNews := new(entity.News)
	if err := r.newsMapper.MapMultilingualFromModel(news, dbNews); err != nil {
		return fmt.Errorf("failed to map news: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Model(&entity.News{
			Model: gorm.Model{
				ID: uint(id),
			},
		}).
		Updates(&entity.News{
			Title:       dbNews.Title,
			Description: dbNews.Description,
			PreviewUrl:  dbNews.PreviewUrl,
			RedirectUrl: dbNews.RedirectUrl,
		}).Error; err != nil {
		return fmt.Errorf("failed to update news: %w", err)
	}

	return nil
}

func (r *NewsRepository) DeleteNews(ctx context.Context, id types.NewsId, force bool) error {
	dbNews := new(entity.News)
	dbNews.ID = uint(id)

	tx := r.db
	if force {
		tx = tx.Unscoped()
	}

	if err := tx.Delete(&dbNews).Error; err != nil {
		return fmt.Errorf("failed to delete news: %w", err)
	}

	return nil
}
