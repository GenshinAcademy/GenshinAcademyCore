package errors

import (
    "fmt"
    "ga/pkg/genshin_core/models"
)

type characterNotFoundError struct{
    id models.ModelId
}

func(err characterNotFoundError) Error() string {
    return fmt.Sprintf("Character with id {%s} not found.", err.id)
}

func CharacterNotFound(id models.ModelId) error {
    return characterNotFoundError{
        id: id,
    }
}


type emptyIdError struct{}

func EmptyId() error {
    return emptyIdError{}
}

func (err emptyIdError) Error() string {
    return fmt.Sprintf("Empty Id provided.")
}