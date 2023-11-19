package mapper

import (
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"
)

type TableMapper struct {
}

func NewTableMapper() *TableMapper {
	return &TableMapper{}
}

func (m *TableMapper) MapFromEntity(input *entity.Table, output *models.Table, language types.Language) error {
	output.Id = types.TableId(input.ID)
	output.Title = input.Title[language]
	output.Description = input.Description[language]
	output.IconUrl = input.IconUrl
	output.RedirectUrl = input.RedirectUrl[language]

	return nil
}

func (m *TableMapper) MapMultilingualFromEntity(input *entity.Table, output *models.TableMultilingual) error {
	output.Title = input.Title
	output.Description = input.Description
	output.IconUrl = input.IconUrl
	output.RedirectUrl = input.RedirectUrl

	return nil
}

func (m *TableMapper) MapMultilingualFromModel(input *models.TableMultilingual, output *entity.Table) error {
	output.Title = input.Title
	output.Description = input.Description
	output.IconUrl = input.IconUrl
	output.RedirectUrl = input.RedirectUrl

	return nil
}
