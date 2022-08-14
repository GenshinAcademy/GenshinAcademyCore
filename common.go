package common

/*Entity types*/
const (
	CharacterType             EntityType = iota
	WeaponType                EntityType = iota
	ArtifactType              EntityType = iota
	CharacterStatType         EntityType = iota
	WeaponMainStatType        EntityType = iota
	WeaponSecondaryStatType   EntityType = iota
	ArtifactMainStatType      EntityType = iota
	ArtifactSecondaryStatType EntityType = iota
	ModifierType              EntityType = iota
)

type EntityId uint64
type EntityType uint32

// Interface for entity
type EntityInterface interface {
	Entity
}

// Base entity
type Entity struct {
	Id   EntityId   `json:"id"`
	Type EntityType `json:"type"`
}

// Compares entities
func (src *Entity) Equals(outer *Entity) bool {
	if src != nil && outer == nil {
		return false
	}
	return src.Type == outer.Type && src.Id == outer.Id
}
