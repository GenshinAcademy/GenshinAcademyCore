package appraiser

import (
	"ga/internal/academy_core/models"
	webModels "ga/internal/services/weasel/appraiser/models"
)

// mapCharacter converts academy_core model to weaselAppraiser model
func (service *Service) mapCharacter(input models.Character) (webModels.WeaselAppraiserCharacter, error) {
	var output webModels.WeaselAppraiserCharacter
	output.CharacterId = input.Character.Id
	output.Name = input.Name
	output.Element = input.Element

	url, err := service.core.GetAssetPath(input.Icons[0].Url)
	if err != nil {
		return output, err
	}

	output.IconUrl = url
	output.StatsProfit = input.Profits

	return output, nil
}
