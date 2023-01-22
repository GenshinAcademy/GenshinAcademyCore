package core

import "ga/pkg/core/repositories"

type GetProviderFunc func() repositories.IRepositoryProvider

type GenshinCoreConfiguration struct {
}

type GenshinCore struct {
	providerFunc GetProviderFunc
}

func (core *GenshinCore) SetProviderFunc(fn GetProviderFunc) *GenshinCore {
	core.providerFunc = fn
	return core
}

func defaultGetProvider() repositories.IRepositoryProvider {
	panic("GetProviderFunc not specified!")
}

func CreateGenshinCore(config GenshinCoreConfiguration) *GenshinCore {
	var core = new(GenshinCore)
	core.providerFunc = defaultGetProvider
	return core
}
