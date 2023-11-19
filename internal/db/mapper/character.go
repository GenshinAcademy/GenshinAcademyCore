package mapper

import (
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"
)

type CharacterMapper struct {
}

func NewCharacterMapper() *CharacterMapper {
	return &CharacterMapper{}
}

func (m *CharacterMapper) MapFromEntity(input *entity.Character, output *models.Character, language types.Language) error {
	output.Id = input.Id
	output.Name = input.Name[language]
	output.Description = input.Description[language]
	output.Element = input.Element
	output.Rarity = input.Rarity
	output.WeaponType = input.WeaponType
	output.IconsUrl = input.Icons.Url

	return nil
}

func (m *CharacterMapper) MapMultilingualFromEntity(input *entity.Character, output *models.CharacterMultilingual) error {
	output.Id = input.Id
	output.Name = input.Name
	output.Description = input.Description
	output.Element = input.Element
	output.Rarity = input.Rarity
	output.WeaponType = input.WeaponType
	output.IconsUrl = input.Icons.Url

	return nil
}

func (m *CharacterMapper) MapMultilingualFromModel(input *models.CharacterMultilingual, output *entity.Character) error {
	output.Id = input.Id
	output.Name = input.Name
	output.Description = input.Description
	output.Element = input.Element
	output.Rarity = input.Rarity
	output.WeaponType = input.WeaponType
	output.Icons.CharacterId = input.Id
	output.Icons.Url = input.IconsUrl

	return nil
}

func (m *CharacterMapper) MapWeaselAppraiserCharacterFromEntity(input *entity.Character, output *models.WeaselAppraiserCharacter, language types.Language) error {
	if err := m.MapFromEntity(input, &output.Character, language); err != nil {
		return err
	}

	output.CharacterArtifactProfits = input.ArtifactProfits.Profits

	return nil
}
