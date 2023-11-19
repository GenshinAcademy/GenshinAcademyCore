package characters

import (
	"context"
	"fmt"
	"ga/internal/models"
	"ga/internal/types"
	"ga/pkg/url"
	"strings"
)

type CharacterRepository interface {
	FindAllByLanguageWithIcons(ctx context.Context, language types.Language, offset int, limit int) ([]models.Character, error)
	FindAllMultilingual(ctx context.Context, offset int, limit int) ([]models.CharacterMultilingual, error)
	CreateCharacter(ctx context.Context, characterMultilingual *models.CharacterMultilingual) error
}

type AssetsService interface {
	GetAssetUrl(assetPath string) (url.Url, error)
	BuildAssetPath(assetType types.AssetType, fileName string) string
}

type Service struct {
	assetsService       AssetsService
	characterRepository CharacterRepository
}

func New(
	assetsService AssetsService,
	characterRepository CharacterRepository,
) *Service {
	return &Service{
		assetsService:       assetsService,
		characterRepository: characterRepository,
	}
}

func (s *Service) GetCharacters(language types.Language, offset int, limit int) ([]models.Character, error) {
	result, err := s.characterRepository.FindAllByLanguageWithIcons(context.Background(), language, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to find all characters: %w", err)
	}
	for i := range result {
		character := &result[i]
		for k, v := range character.IconsUrl {
			if !url.IsUrl(v) && v != "" {
				iconURL, err := s.assetsService.GetAssetUrl(v)
				if err != nil {
					return nil, fmt.Errorf("failed to get asset path: %w", err)
				}
				character.IconsUrl[k] = string(iconURL)
			}
		}
	}
	return result, nil
}

func (s *Service) CreateCharacter(character *models.CharacterMultilingual) error {
	if character.Id == "" {
		defaultName, ok := character.Name[types.DefaultLanguage]
		if !ok {
			return fmt.Errorf("default name was not set")
		}

		character.Id = types.CharacterId(strings.ToLower(strings.ReplaceAll(defaultName, " ", "_")))
	}
	if character.IconsUrl == nil {
		character.IconsUrl = map[types.IconType]string{
			types.FrontFace: string(character.Id),
		}
	}

	for k, v := range character.IconsUrl {
		if !url.IsUrl(v) && v != "" {
			character.IconsUrl[k] = s.assetsService.BuildAssetPath(types.CharactersIconsAsset, v)
		}
	}

	if err := s.characterRepository.CreateCharacter(context.Background(), character); err != nil {
		return fmt.Errorf("failed to create multilingual character: %w", err)
	}

	return nil
}
