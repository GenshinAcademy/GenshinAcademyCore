package service

import (
	models_db "genshinacademycore/models/db"
	models_web "genshinacademycore/models/web"
	"genshinacademycore/repository"
)

type FerretServiceInterface interface {
	GetAllCharactersStats() (*[]models_web.CharacterArtifactStatsProfit, error)
	// GetAllCharacters() (*[]models.Character, error)
	// GetCharacter(id string) (*[]models.Character, error)
}

type FerretService struct {
	FerretRepository repository.FerretRepositoryInterface
}

func NewFerretService(repoFerret repository.FerretRepositoryInterface) FerretServiceInterface {
	return &FerretService{
		FerretRepository: repoFerret,
	}
}

func (f *FerretService) GetAllCharactersStats() (*[]models_web.CharacterArtifactStatsProfit, error) {
	dbCharacters, err := f.FerretRepository.GetAllCharactersStats()
	if err != nil {
		return nil, err
	}

	result := []models_web.CharacterArtifactStatsProfit{}
	for _, element := range *dbCharacters {
		var stats *models_db.ArtifactStats = (*models_db.ArtifactStats)(&element.StatsProfit)
		result = append(result, models_web.CharacterArtifactStatsProfit{
			Character: CharacterFromDBModel(&element),
			StatsProfit: models_web.StatsProfit{
				Flower:   FlowerFromDBModel(stats),
				Feather:  FeatherFromDBModel(stats),
				Sands:    SandsFromDBModel(stats),
				Goblet:   GobletFromDBModel(stats),
				Circlet:  CircletFromDBModel(stats),
				Substats: SubstatsFromDBModel(stats),
			},
		})
	}

	return &result, nil
	//return f.FerretRepository.GetAllCharactersStats()
}

func CharacterFromDBModel(character *models_db.Character) models_web.Character {
	return models_web.Character{
		ID:      character.ID,
		Name:    character.Name.English,
		Element: character.Element,
	}
}

func FlowerFromDBModel(stats *models_db.ArtifactStats) models_web.Flower {
	return models_web.Flower{
		Health: stats.Health,
	}
}

func FeatherFromDBModel(stats *models_db.ArtifactStats) models_web.Feather {
	return models_web.Feather{
		Attack: stats.Attack,
	}
}

func SandsFromDBModel(stats *models_db.ArtifactStats) models_web.Sands {
	return models_web.Sands{
		AttackPercentage:  stats.AttackPercentage,
		HealthPercentage:  stats.HealthPercentage,
		DefensePercentage: stats.DefensePercentage,
		ElementalMastery:  stats.ElementalMastery,
		EnergyRecharge:    stats.EnergyRecharge,
	}
}

func GobletFromDBModel(stats *models_db.ArtifactStats) models_web.Goblet {
	return models_web.Goblet{
		AttackPercentage:  stats.AttackPercentage,
		HealthPercentage:  stats.HealthPercentage,
		DefensePercentage: stats.DefensePercentage,
		ElementalMastery:  stats.ElementalMastery,
		PhysicalDamage:    stats.PhysicalDamage,
		ElementalDamage:   stats.ElementalDamage,
	}
}

func CircletFromDBModel(stats *models_db.ArtifactStats) models_web.Circlet {
	return models_web.Circlet{
		AttackPercentage:  stats.AttackPercentage,
		HealthPercentage:  stats.HealthPercentage,
		DefensePercentage: stats.DefensePercentage,
		ElementalMastery:  stats.ElementalMastery,
		CritRate:          stats.CritRate,
		CritDamage:        stats.CritDamage,
		Heal:              stats.Heal,
	}
}

func SubstatsFromDBModel(stats *models_db.ArtifactStats) models_web.Substats {
	return models_web.Substats{
		Attack:            stats.Attack,
		AttackPercentage:  stats.AttackPercentage,
		Health:            stats.Health,
		HealthPercentage:  stats.HealthPercentage,
		Defense:           stats.Defense,
		DefensePercentage: stats.DefensePercentage,
		ElementalMastery:  stats.ElementalMastery,
		EnergyRecharge:    stats.EnergyRecharge,
		CritDamage:        stats.CritDamage,
		CritRate:          stats.CritRate,
	}
}

// func (f *FerretService) GetAllCharacters() (*[]models.Character, error) {
// 	return f.FerretRepository.GetAllCharacters()
// }

// func (f *FerretService) GetCharacter(id string) (*[]models.Character, error) {
// 	return f.FerretRepository.GetCharacter(id)
// }
