package ferret

import (
	"ga/internal/academy_core"
	"ga/internal/ferret/web_models"
)

type FerretService struct {
	core academy_core.AcademyCore
}

func CreateService(core academy_core.AcademyCore) *FerretService {
	var result *FerretService = new(FerretService)
	return result
}

func (service *FerretService) GetAllCharactersWithProfits() []web_models.FerretCharacter {
	panic("Not implemented")
}
