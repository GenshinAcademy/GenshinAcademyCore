package find_parameters

import (
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/enums"
)

type WeaponFindParameters struct {
	FindParameters
	Types []enums.WeaponType
}

func (param WeaponFindParameters) AddType(weaponType enums.WeaponType) WeaponFindParameters {
	param.Types = append(param.Types, weaponType)
	return param
}

func FindByWeaponId(weaponId string) WeaponFindParameters {
	return WeaponFindParameters{
		FindParameters: FindParameters{
			Ids:          []models.ModelId{models.ModelId(weaponId)},
			SliceOptions: SliceParameters{Offset: 0, Limit: 1},
		},
	}
}
