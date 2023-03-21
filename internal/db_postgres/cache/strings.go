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

type NewsStrings struct {
	Title       db_models.DBKey
	Description db_models.DBKey
}

func (cache *Cache) GetCharacterStrings(key db_models.DBKey) *CharacterStrings {
	var val, ok = cache.characterStrings[key]
    if !ok {
        cache.Lock()
        val = new(CharacterStrings)
        cache.characterStrings[key] = val
        cache.Unlock()
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

func (cache *Cache) GetNewsStrings(key db_models.DBKey) *NewsStrings {
    var val, ok = cache.newsStrings[key]
    if !ok {
        val = new(NewsStrings)
        cache.Lock()
        cache.newsStrings[key] = val
        cache.Unlock()
    }
    return val
}

func (cache *Cache) UpdateNewsStrings(model *db_models.News) *NewsStrings{
    cache.Lock()
    var strings = cache.GetNewsStrings(model.Id)
    if strings == nil {
        strings = new(NewsStrings)
        cache.newsStrings[model.Id] = strings
    }

    strings.Title = model.TitleId
    strings.Description = model.DescriptionId
    cache.Unlock()

    return strings
}
