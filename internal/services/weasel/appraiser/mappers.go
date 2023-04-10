package appraiser

import (
	"ga/internal/academy_core/models"
	webModels "ga/internal/services/weasel/appraiser/models"
)

// mapCharacter converts academy_core model to weaselAppraiser model
func (service *Service) mapCharacter(input models.Character) (webModels.WeaselAppraiserCharacter, error) {
	var output webModels.WeaselAppraiserCharacter
	output.CharacterId = string(input.Character.Id)
	output.Name = input.Name
	output.Element = uint8(input.Element)

	url, err := service.core.GetAssetPath(input.Icons[0].Url)
	if err != nil {
		return output, err
	}
	output.IconUrl = url

	output.StatsProfit = make([]webModels.StatsProfit, len(input.Profits))

	for i, stat := range input.Profits {
		var statProfit = &output.StatsProfit[i]
		statProfit.Slot = string(stat.Slot)
		statProfit.Attack = int(stat.Attack)
		statProfit.AttackPercentage = int(stat.AttackPercentage)
		statProfit.CritDamage = int(stat.CritDamage)
		statProfit.CritRate = int(stat.CritRate)
		statProfit.Defense = int(stat.Defense)
		statProfit.DefensePercentage = int(stat.DefensePercentage)
		statProfit.ElementalDamage = int(stat.ElementalDamage)
		statProfit.ElementalMastery = int(stat.ElementalMastery)
		statProfit.EnergyRecharge = int(stat.EnergyRecharge)
		statProfit.Heal = int(stat.Heal)
		statProfit.Health = int(stat.Health)
		statProfit.HealthPercentage = int(stat.HealthPercentage)
		statProfit.PhysicalDamage = int(stat.PhysicalDamage)
	}

	return output, nil
}
