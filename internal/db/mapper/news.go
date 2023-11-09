package mapper

import (
	"ga/internal/db/entity"
	"ga/internal/models"
	"ga/internal/types"
)

type NewsMapper struct {
}

func NewNewsMapper() *NewsMapper {
	return &NewsMapper{}
}

func (m *NewsMapper) MapFromEntity(input *entity.News, output *models.News, language types.Language) error {
	output.Id = types.NewsId(input.ID)
	output.Title = input.Title[language]
	output.Description = input.Description[language]
	output.Preview = input.PreviewUrl[language]
	output.RedirectUrl = input.RedirectUrl[language]
	output.CreatedAt = input.CreatedAt

	return nil
}

func (m *NewsMapper) MapMultilingualFromEntity(input *entity.News, output *models.NewsMultilingual) error {
	output.Title = input.Title
	output.Description = input.Description
	output.Preview = input.PreviewUrl
	output.RedirectUrl = input.RedirectUrl
	output.CreatedAt = input.CreatedAt

	return nil
}

func (m *NewsMapper) MapMultilingualFromModel(input *models.NewsMultilingual, output *entity.News) error {
	output.Title = input.Title
	output.Description = input.Description
	output.PreviewUrl = input.Preview
	output.RedirectUrl = input.RedirectUrl
	output.CreatedAt = input.CreatedAt

	return nil
}
