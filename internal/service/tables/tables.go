package tables

import (
	"context"
	"fmt"
	"ga/internal/models"
	"ga/internal/types"
	"ga/pkg/url"
)

type TablesRepository interface {
	FindAllTables(ctx context.Context, language types.Language, offset int, limit int, sort string) ([]models.Table, error)
	CreateTable(ctx context.Context, table *models.TableMultilingual) error
	UpdateTable(ctx context.Context, id types.TableId, Tables *models.TableMultilingual) error
	DeleteTable(ctx context.Context, id types.TableId, force bool) error
}

type AssetsService interface {
	GetAssetUrl(assetPath string) (url.Url, error)
	BuildAssetPath(assetType types.AssetType, fileName string) string
}

type Service struct {
	assetsService    AssetsService
	tablesRepository TablesRepository
}

func New(
	assetService AssetsService,
	tablesRepository TablesRepository,
) *Service {
	return &Service{
		assetsService:    assetService,
		tablesRepository: tablesRepository,
	}
}
func (s *Service) GetTables(language types.Language, offset int, limit int, sort string) ([]models.Table, error) {
	result, err := s.tablesRepository.FindAllTables(context.Background(), language, offset, limit, sort)
	if err != nil {
		return nil, fmt.Errorf("failed to get tables: %w", err)
	}

	for i := range result {
		table := &result[i]
		if !url.IsUrl(table.IconUrl) && table.IconUrl != "" {
			iconURL, err := s.assetsService.GetAssetUrl(table.IconUrl)
			if err != nil {
				return nil, fmt.Errorf("failed to get asset path: %w", err)
			}
			table.IconUrl = string(iconURL)
		}
	}

	return result, nil
}

func (s *Service) CreateTable(table *models.TableMultilingual) error {
	if !url.IsUrl(table.IconUrl) && table.IconUrl != "" {
		table.IconUrl = s.assetsService.BuildAssetPath(types.TablesAsset, table.IconUrl)
	}

	if err := s.tablesRepository.CreateTable(context.Background(), table); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (s *Service) UpdateTable(id types.TableId, table *models.TableMultilingual) error {
	if !url.IsUrl(table.IconUrl) && table.IconUrl != "" {
		table.IconUrl = s.assetsService.BuildAssetPath(types.TablesAsset, table.IconUrl)
	}

	if err := s.tablesRepository.UpdateTable(context.Background(), id, table); err != nil {
		return fmt.Errorf("failed to update table: %w", err)
	}

	return nil
}

func (s *Service) DeleteTable(id types.TableId, force bool) error {
	if err := s.tablesRepository.DeleteTable(context.Background(), id, force); err != nil {
		return fmt.Errorf("failed to delete table: %w", err)
	}

	return nil
}
