package ferret

import (
	"ga/internal/academy_core/models"
	"ga/internal/ferret/web_models"
)

// mapCharacter converts academy_core model to ferret model
func (service *FerretService) mapCharacter(input models.Character) web_models.FerretCharacter {
	var output web_models.FerretCharacter
	output.CharacterId = string(input.Character.Id)
	output.Name = input.Name
	output.Element = uint8(input.Element)
	output.IconUrl = input.Icons[0].Url
	output.StatsProfit = []web_models.StatsProfit{}
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

	return output
}
