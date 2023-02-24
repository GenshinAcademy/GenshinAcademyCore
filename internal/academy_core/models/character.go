package models

import (
	artifact_proft "ga/internal/academy_core/value_objects/artifact_profit"
	"ga/pkg/genshin_core/models"
)


type Character struct {
	AcademyModel
	models.Character
	Profits []artifact_proft.ArtifactProfit
}
