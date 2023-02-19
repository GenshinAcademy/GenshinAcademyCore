package db_models

type DbArtifactProfit struct {
	CharacterId       DBKey `gorm:"primaryKey"`
	SlotId            DBKey `gorm:"primaryKey"`
	SlotName          string
	Attack            uint16
	AttackPercentage  uint16
	Health            uint16
	HealthPercentage  uint16
	Defense           uint16
	DefensePercentage uint16
	ElementalMastery  uint16
	EnergyRecharge    uint16
	ElementalDamage   uint16
	CritRate          uint16
	CritDamage        uint16
	PhysicalDamage    uint16
	Heal              uint16
}
