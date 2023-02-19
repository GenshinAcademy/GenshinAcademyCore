package genshin

import (
    academy_find_parameters "ga/internal/academy_core/repositories/find_parameters"
    "ga/internal/db_postgres/repositories/academy"
    "ga/pkg/genshin_core/languages"
    "ga/pkg/genshin_core/models"
    find_parameters "ga/pkg/genshin_core/repositories/find_parameters"
)

type GenshinCharacterRepository struct {
    academyRepo *academy.PostgresCharacterRepository
}

func CreateGenshinCharacterRepository(repo *academy.PostgresCharacterRepository) *GenshinCharacterRepository {
    var genshinRepo = new(GenshinCharacterRepository)
    genshinRepo.academyRepo = repo

    return genshinRepo
}
func(repo GenshinCharacterRepository) GetLanguage() languages.Language {
    return languages.Language(repo.academyRepo.GetLanguage().LanguageName)
}

func(repo GenshinCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []models.ModelId {
    var academyParams = academy_find_parameters.CharacterFindParameters{
        CharacterFindParameters: parameters,
    }

    return repo.academyRepo.GetCharacterIds(academyParams)
}
