package academy_core

import (
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"

	"ga/pkg/genshin_core"
	"ga/pkg/genshin_core/models/languages"
)

type GetAcademyProviderFunc func(models.Language) repositories.IRepositoryProvider
type GetLanguageRepositoryFunc func() repositories.ILanguageRepository

type AcademyCoreConfiguration struct {
	genshin_core.GenshinCoreConfiguration
	ProviderFunc     GetAcademyProviderFunc
	LanguageRepoFunc GetLanguageRepositoryFunc
}

type AcademyCore struct {
	getLanguageRepository GetLanguageRepositoryFunc
	getProvider           GetAcademyProviderFunc
	genshinCore           *genshin_core.GenshinCore
}

func CreateAcademyCore(configuration AcademyCoreConfiguration) *AcademyCore {
	var core = new(AcademyCore)
	core.genshinCore = genshin_core.CreateGenshinCore(configuration.GenshinCoreConfiguration)

	core.getLanguageRepository = defaultGetLanguageRepository
	core.getProvider = defaultGetProvider

	if configuration.LanguageRepoFunc != nil {
		core.getLanguageRepository = configuration.LanguageRepoFunc
	}
	if configuration.ProviderFunc != nil {
		core.getProvider = configuration.ProviderFunc
	}

	return core
}

func defaultGetLanguageRepository() repositories.ILanguageRepository {
	panic("GetLanguagerepositoryFunc not specified!")
}

func defaultGetProvider(models.Language) repositories.IRepositoryProvider {
	panic("GetProviderFunc not specified!")
}

func (core *AcademyCore) AsGenshinCore() *genshin_core.GenshinCore {
	return core.genshinCore
}

func (core *AcademyCore) GetDefaultLanguage() models.Language {
	var language = core.getLanguageRepository().FindLanguage(core.genshinCore.GetDefaultLanguageName())
	return language
}

// GetLanguageRepository returns an object to operate with languages.
func (core *AcademyCore) GetLanguageRepository() repositories.ILanguageRepository {
	return core.getLanguageRepository()
}

func (core *AcademyCore) GetDefaultProvider() repositories.IRepositoryProvider {
	var language = core.getLanguageRepository().FindLanguage(core.genshinCore.GetDefaultLanguageName())
	return core.getProvider(language)
}

func (core *AcademyCore) GetProvider(languageName languages.Language) repositories.IRepositoryProvider {
	var language = core.getLanguageRepository().FindLanguage(languageName)
	return core.getProvider(language)
}
