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

type TableStrings struct {
    Title       db_models.DBKey
    Description db_models.DBKey
}

func (cache *Cache) GetCharacterStrings(key db_models.DBKey) *CharacterStrings {
	var val, ok = cache.characterStrings[key]
    if !ok {
        val = new(CharacterStrings)
        cache.Lock()
        cache.characterStrings[key] = val
        cache.Unlock()
	}
	return val
}

func (cache *Cache) UpdateCharacterStrings(model *db_models.Character) *CharacterStrings {
    var strings = cache.GetCharacterStrings(model.Id)

	cache.Lock()
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
    var strings = cache.GetNewsStrings(model.Id)

    cache.Lock()
    if strings == nil {
        strings = new(NewsStrings)
        cache.newsStrings[model.Id] = strings
    }

    strings.Title = model.TitleId
    strings.Description = model.DescriptionId
    cache.Unlock()

    return strings
}

func (cache *Cache) GetTableStrings(key db_models.DBKey) *TableStrings {
    var val, ok = cache.tableStrings[key]
    if !ok {
        val = new(TableStrings)
        cache.Lock()
        cache.tableStrings[key] = val
        cache.Unlock()
    }
    return val
}

func (cache *Cache) UpdateTableStrings(model *db_models.Table) *TableStrings{
    var strings = cache.GetTableStrings(model.Id)

    cache.Lock()
    if strings == nil {
        strings = new(TableStrings)
        cache.tableStrings[model.Id] = strings
    }

    strings.Title = model.TitleId
    strings.Description = model.DescriptionId
    cache.Unlock()

    return strings
}
