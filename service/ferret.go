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
	for _, character := range *dbCharacters {
		var web_stats models_web.StatsProfit
		for _, stats := range character.StatsProfit {
			switch stats.Slot.Name {
			case "Flower":
				web_stats.Flower = StatsFromDBModel(&stats).(models_web.Flower)
			case "Feather":
				web_stats.Feather = StatsFromDBModel(&stats).(models_web.Feather)
			case "Sands":
				web_stats.Sands = StatsFromDBModel(&stats).(models_web.Sands)
			case "Goblet":
				web_stats.Goblet = StatsFromDBModel(&stats).(models_web.Goblet)
			case "Circlet":
				web_stats.Circlet = StatsFromDBModel(&stats).(models_web.Circlet)
			default:
				web_stats.Substats = StatsFromDBModel(&stats).(models_web.Substats)
			}
		}

		result = append(result, models_web.CharacterArtifactStatsProfit{
			Character:   CharacterFromDBModel(&character),
			StatsProfit: web_stats,
		})
	}
	return &result, nil
}

func CharacterFromDBModel(character *models_db.Character) models_web.Character {
	return models_web.Character{
		ID:      character.ID,
		Name:    character.Name.English,
		Element: character.Element.Name,
		IconURL: character.IconURL,
	}
}

func StatsFromDBModel(stats *models_db.StatsProfit) interface{} {
	switch stats.Slot.Name {
	case "Flower":
		return models_web.Flower{
			Health: stats.Health,
		}
	case "Feather":
		return models_web.Feather{
			Attack: stats.Attack,
		}
	case "Sands":
		return models_web.Sands{
			AttackPercentage:  stats.AttackPercentage,
			HealthPercentage:  stats.HealthPercentage,
			DefensePercentage: stats.DefensePercentage,
			ElementalMastery:  stats.ElementalMastery,
			EnergyRecharge:    stats.EnergyRecharge,
		}
	case "Goblet":
		return models_web.Goblet{
			AttackPercentage:  stats.AttackPercentage,
			HealthPercentage:  stats.HealthPercentage,
			DefensePercentage: stats.DefensePercentage,
			ElementalMastery:  stats.ElementalMastery,
			PhysicalDamage:    stats.PhysicalDamage,
			ElementalDamage:   stats.ElementalDamage,
		}
	case "Circlet":
		return models_web.Circlet{
			AttackPercentage:  stats.AttackPercentage,
			HealthPercentage:  stats.HealthPercentage,
			DefensePercentage: stats.DefensePercentage,
			ElementalMastery:  stats.ElementalMastery,
			CritRate:          stats.CritRate,
			CritDamage:        stats.CritDamage,
			Heal:              stats.Heal,
		}
	default:
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
}
