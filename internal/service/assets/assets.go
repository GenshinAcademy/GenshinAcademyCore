package assets

import (
	"ga/internal/types"
	"ga/pkg/url"
	"path/filepath"
)

var validAssetTypes = map[types.AssetType]bool{
	types.CharactersAsset:      true,
	types.CharactersIconsAsset: true,
	types.TablesAsset:          true,
	types.NewsAsset:            true,
	types.OpenGraphAsset:       true,
}

type Service struct {
	assetsPath     string
	assetsHost     string
	assetTypePaths map[types.AssetType]string
}

func New(
	assetsPath string,
	assetsHost string,
) *Service {
	assetTypePaths := make(map[types.AssetType]string, len(validAssetTypes))
	for assetType := range validAssetTypes {
		assetTypePaths[assetType] = filepath.Join(assetsPath, string(assetType))
	}

	return &Service{
		assetsPath:     assetsPath,
		assetsHost:     assetsHost,
		assetTypePaths: assetTypePaths,
	}
}

func (s *Service) IsValidAssetType(assetType string) bool {
	v, ok := validAssetTypes[types.AssetType(assetType)]
	if ok && v {
		return true
	}

	return false
}

func (s *Service) GetPossibleAssetTypes() (result []string) {
	result = make([]string, 0)
	for k, v := range validAssetTypes {
		if v {
			result = append(result, string(k))
		}
	}

	return result
}

func (s *Service) GetPathForAssetType(assetType types.AssetType) string {
	return filepath.Join(s.assetsPath, string(assetType))
}

func (s *Service) GetAssetUrl(assetPath string) (url.Url, error) {
	return url.CreateUrl(s.assetsHost, assetPath)
}

func (s *Service) BuildAssetPath(assetType types.AssetType, fileName string) string {
	if path, ok := s.assetTypePaths[assetType]; ok {
		return filepath.Join(path, fileName)
	}

	return ""
}
