package news

import (
	"context"
	"fmt"
	"ga/internal/models"
	"ga/internal/types"
	"ga/pkg/url"
)

type NewsRepository interface {
	FindAllNews(ctx context.Context, language types.Language, offset int, limit int, sort string) ([]models.News, error)
	CreateNews(ctx context.Context, news *models.NewsMultilingual) error
	UpdateNews(ctx context.Context, id types.NewsId, news *models.NewsMultilingual) error
}

type AssetsService interface {
	GetAssetUrl(assetPath string) (url.Url, error)
	BuildAssetPath(assetType types.AssetType, fileName string) string
}

type Service struct {
	assetsService  AssetsService
	newsRepository NewsRepository
}

func New(
	assetsService AssetsService,
	newsRepository NewsRepository,
) *Service {
	return &Service{
		assetsService:  assetsService,
		newsRepository: newsRepository,
	}
}

func (s *Service) GetNews(language types.Language, offset int, limit int, sort string) ([]models.News, error) {
	result, err := s.newsRepository.FindAllNews(context.Background(), language, offset, limit, sort)
	if err != nil {
		return nil, fmt.Errorf("failed to get news: %w", err)
	}

	for i := range result {
		news := &result[i]
		if !url.IsUrl(news.Preview) && news.Preview != "" {
			iconURL, err := s.assetsService.GetAssetUrl(news.Preview)
			if err != nil {
				return nil, fmt.Errorf("failed to get asset path: %w", err)
			}
			news.Preview = string(iconURL)
		}
	}

	return result, nil
}

func (s *Service) CreateNews(news *models.NewsMultilingual) error {
	for lang, fileName := range news.Preview {
		if !url.IsUrl(fileName) && fileName != "" {
			news.Preview[lang] = s.assetsService.BuildAssetPath(types.NewsAsset, fileName)
		}
	}

	if err := s.newsRepository.CreateNews(context.Background(), news); err != nil {
		return fmt.Errorf("failed to create news: %w", err)
	}

	return nil
}

func (s *Service) UpdateNews(id types.NewsId, news *models.NewsMultilingual) error {
	for lang, fileName := range news.Preview {
		if !url.IsUrl(fileName) && fileName != "" {
			news.Preview[lang] = s.assetsService.BuildAssetPath(types.NewsAsset, fileName)
		}
	}

	if err := s.newsRepository.UpdateNews(context.Background(), id, news); err != nil {
		return fmt.Errorf("failed to update news: %w", err)
	}

	return nil
}
