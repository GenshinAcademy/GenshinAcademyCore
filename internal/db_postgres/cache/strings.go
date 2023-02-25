package cache

import (
	db_models "ga/internal/db_postgres/models"
)

type CharacterStrings struct {
	Name        db_models.DBKey
	FullName    db_models.DBKey
	Description db_models.DBKey
	Title       db_models.DBKey
}

func (cache *Cache) GetCharacterStrings(key db_models.DBKey) *CharacterStrings {
	var val, ok = cache.characterStrings[key]
	if !ok {
		return nil
	}
	return val
}

func (cache *Cache) UpdateCharacterStrings(model *db_models.Character) *CharacterStrings {
	cache.Lock()
	var strings = cache.GetCharacterStrings(model.Id)
	if strings == nil {
		strings = new(CharacterStrings)
		cache.characterStrings[model.Id] = strings
	}

	strings.Name = model.NameId
	strings.FullName = model.FullNameId
	strings.Description = model.DescriptionId
	strings.Title = model.TitleId
	cache.Unlock()

	return strings
}
