package db_models

// DbArtifactProfit represents artifact stats profit for Genshin characters in database.
type DbArtifactProfit struct {
	CharacterId       DBKey  `gorm:"primaryKey"`
	SlotId            DBKey  `gorm:"primaryKey"`
	Attack            uint16 `gorm:"not null"`
	AttackPercentage  uint16 `gorm:"not null"`
	Health            uint16 `gorm:"not null"`
	HealthPercentage  uint16 `gorm:"not null"`
	Defense           uint16 `gorm:"not null"`
	DefensePercentage uint16 `gorm:"not null"`
	ElementalMastery  uint16 `gorm:"not null"`
	EnergyRecharge    uint16 `gorm:"not null"`
	ElementalDamage   uint16 `gorm:"not null"`
	CritRate          uint16 `gorm:"not null"`
	CritDamage        uint16 `gorm:"not null"`
	PhysicalDamage    uint16 `gorm:"not null"`
	Heal              uint16 `gorm:"not null"`
}
