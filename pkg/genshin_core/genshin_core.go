package genshin_core

import (
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories"
)

type GetProviderFunc func(*languages.Language) repositories.RepositoryProvider

type GenshinCoreConfiguration struct {
	DefaultLanguage languages.Language
	ProviderFunc    GetProviderFunc
}

type GenshinCore struct {
	providerFunc        GetProviderFunc
	defaultLanguageName languages.Language
}

func (core *GenshinCore) GetDefaultLanguageName() *languages.Language {
	return &core.defaultLanguageName
}

func defaultGetProvider(*languages.Language) repositories.RepositoryProvider {
	panic("GetProviderFunc not specified!")
}
func CreateGenshinCore(config GenshinCoreConfiguration) *GenshinCore {
	var core = new(GenshinCore)
	core.defaultLanguageName = config.DefaultLanguage
	core.providerFunc = defaultGetProvider

	if config.ProviderFunc != nil {
		core.providerFunc = config.ProviderFunc
	}
	return core
}

func (core *GenshinCore) GetDefaultProvider() repositories.RepositoryProvider {
	return core.GetProvider(&core.defaultLanguageName)
}

func (core *GenshinCore) GetProvider(language *languages.Language) repositories.RepositoryProvider {
	return core.providerFunc(language)
}
