package models

import (
	artifact_profit "ga/internal/academy_core/value_objects/artifact_profit"
	"ga/pkg/genshin_core/models"
)

type Character struct {
	AcademyModel
	models.Character
	Profits []artifact_profit.ArtifactProfit `extensions:"x-order=5"`
} //@name AcademyCharacter
