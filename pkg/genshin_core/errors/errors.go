package errors

import (
	"fmt"
	"ga/pkg/genshin_core/models"
)

type characterNotFoundError struct {
	id models.ModelId
}

type weaponNotFoundError struct {
	id models.ModelId
}

func (err characterNotFoundError) Error() string {
	return fmt.Sprintf("Character with id {%s} not found.", err.id)
}

func (err weaponNotFoundError) Error() string {
	return fmt.Sprintf("Weapon with id {%s} not found.", err.id)
}

// CharacterNotFound Indicates that character was not found
func CharacterNotFound(id models.ModelId) error {
	return characterNotFoundError{
		id: id,
	}
}

// WeaponNotFound Indicates that character was not found
func WeaponNotFound(id models.ModelId) error {
	return weaponNotFoundError{
		id: id,
	}
}

type emptyIdError struct{}

// EmptyId Indicates that Id was empty.
func EmptyId() error {
	return emptyIdError{}
}

func (err emptyIdError) Error() string {
	return "Empty Id provided."
}
