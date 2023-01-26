package core

import (
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"
)

type GetProviderFunc func(models.Language) repositories.IRepositoryProvider
type GetLanguageRepoFunc func() repositories.ILanguageRepository

type GenshinCoreConfiguration struct {
	DefaultLanguage  string
	ProviderFunc     GetProviderFunc
	LanguageRepoFunc GetLanguageRepoFunc
}

type GenshinCore struct {
	providerFunc        GetProviderFunc
	languageRepoFunc    GetLanguageRepoFunc
	defaultLanguageName string
}

func (core *GenshinCore) GetDefaultLanguageName() string {
	return core.defaultLanguageName
}

func defaultGetProvider(language models.Language) repositories.IRepositoryProvider {
	panic("GetProviderFunc not specified!")
}

func defaultGetLanguageRepo() repositories.ILanguageRepository {
	panic("GetLanguageRepoFunc not specified!")
}

func CreateGenshinCore(config GenshinCoreConfiguration) *GenshinCore {
	var core = new(GenshinCore)
	core.defaultLanguageName = config.DefaultLanguage
	core.providerFunc = defaultGetProvider
	core.languageRepoFunc = defaultGetLanguageRepo

	if config.LanguageRepoFunc != nil {
		core.languageRepoFunc = config.LanguageRepoFunc
	}
	if config.LanguageRepoFunc != nil {
		core.providerFunc = config.ProviderFunc
	}
	return core
}

func (core *GenshinCore) GetLanguageRepository() repositories.ILanguageRepository {
	return core.languageRepoFunc()
}

func (core *GenshinCore) GetDefaultProvider() repositories.IRepositoryProvider {
	var defaultLanguage = core.languageRepoFunc().FindLanguage(core.defaultLanguageName)
	return core.GetProvider(defaultLanguage)
}

func (core *GenshinCore) GetProvider(language models.Language) repositories.IRepositoryProvider {
	return core.providerFunc(language)
}
