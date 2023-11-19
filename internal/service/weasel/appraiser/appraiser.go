package appraiser

import (
	"context"
	"fmt"
	"ga/internal/models"
	"ga/internal/types"
	"ga/pkg/url"
	log "github.com/sirupsen/logrus"
)

type AssetsService interface {
	GetAssetUrl(assetPath string) (url.Url, error)
}

type CharacterRepository interface {
	FindAllByLanguageWithArtifactProfitsAndIcons(ctx context.Context, language types.Language, offset int, limit int) ([]models.WeaselAppraiserCharacter, error)
}

type ArtifactProfitsRepository interface {
	FindOneByCharacterId(ctx context.Context, id types.CharacterId) (types.CharacterArtifactProfits, error)
	UpdateArtifactProfits(ctx context.Context, id types.CharacterId, artifactProfits types.CharacterArtifactProfits) error
}

type Service struct {
	assetsService             AssetsService
	characterRepository       CharacterRepository
	artifactProfitsRepository ArtifactProfitsRepository
}

func New(
	assetsService AssetsService,
	characterRepository CharacterRepository,
	artifactProfitsRepository ArtifactProfitsRepository,
) *Service {
	return &Service{
		assetsService:             assetsService,
		characterRepository:       characterRepository,
		artifactProfitsRepository: artifactProfitsRepository,
	}
}

func (s *Service) GetAll(language types.Language) ([]models.WeaselAppraiserCharacter, error) {
	result, err := s.characterRepository.FindAllByLanguageWithArtifactProfitsAndIcons(context.Background(), language, 0, 100)
	if err != nil {
		return nil, fmt.Errorf("failed to find all characters: %w", err)
	}
	for i := range result {
		c := &result[i]
		for k, v := range c.IconsUrl {
			if !url.IsUrl(v) && v != "" {
				iconURL, err := s.assetsService.GetAssetUrl(v)
				if err != nil {
					return nil, fmt.Errorf("failed to get asset path: %w", err)
				}
				c.IconsUrl[k] = string(iconURL)
			}
		}
	}

	return result, nil
}

func (s *Service) UpdateStats(id types.CharacterId, artifactProfits types.CharacterArtifactProfits) error {
	srcArtifactProfits, err := s.artifactProfitsRepository.FindOneByCharacterId(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to find artifact profits: %w", err)
	}

	log.Info(srcArtifactProfits)
	if srcArtifactProfits == nil {
		return fmt.Errorf("failed to find artifact profits")
	}

	for slot, stat := range artifactProfits {
		for statType, statProfit := range stat {
			srcArtifactProfits[slot][statType] = statProfit
		}
	}

	if err := s.artifactProfitsRepository.UpdateArtifactProfits(context.Background(), id, srcArtifactProfits); err != nil {
		return fmt.Errorf("failed to update artifact profits: %w", err)
	}

	return nil
}

var artifactSlots = map[types.ArtifactSlot]bool{
	types.Circlet:  true,
	types.Flower:   true,
	types.Goblet:   true,
	types.Plume:    true,
	types.Sands:    true,
	types.SubStats: true,
}

var artifactStats = map[types.StatType]bool{
	types.Attack:            true,
	types.AttackPercentage:  true,
	types.Health:            true,
	types.HealthPercentage:  true,
	types.Defence:           true,
	types.DefencePercentage: true,
	types.ElementalMastery:  true,
	types.EnergyRecharge:    true,
	types.ElementalDamage:   true,
	types.CritRate:          true,
	types.CritDamage:        true,
	types.PhysicalDamage:    true,
	types.Heal:              true,
}

func (s *Service) GetValidArtifactSlots() map[types.ArtifactSlot]bool {
	return artifactSlots
}

func (s *Service) GetValidArtifactTypes() map[types.StatType]bool {
	return artifactStats
}

func (s *Service) ValidateArtifactProfitsStructure(requestData types.CharacterArtifactProfits) (valid bool, invalidArtifactSlots []types.ArtifactSlot, invalidArtifactStats []types.StatType) {
	valid = true

	for artifact := range requestData {
		if _, ok := artifactSlots[artifact]; !ok {
			valid = false
			invalidArtifactSlots = append(invalidArtifactSlots, artifact)
		}
	}

	for _, artifactData := range requestData {
		for stat := range artifactData {
			if _, ok := artifactStats[stat]; !ok {
				valid = false
				invalidArtifactStats = append(invalidArtifactStats, stat)
			}
		}
	}

	return
}
