package service

import (
	models_db "genshinacademycore/models/db"
	"genshinacademycore/repository"
)

type FerretServiceInterface interface {
	GetAllCharactersStats() (*[]models_db.Character, error)
	// GetAllCharacters() (*[]models.Character, error)
	// GetCharacter(id string) (*[]models.Character, error)
}

type FerretService struct {
	FerretRepository repository.FerretRepositoryInterface
}

func NewFerretService(repoFerret repository.FerretRepositoryInterface) FerretServiceInterface {
	return &FerretService{
		FerretRepository: repoFerret,
	}
}

func (f *FerretService) GetAllCharactersStats() (*[]models_db.Character, error) {
	return f.FerretRepository.GetAllCharactersStats()
}

// func (f *FerretService) GetAllCharacters() (*[]models.Character, error) {
// 	return f.FerretRepository.GetAllCharacters()
// }

// func (f *FerretService) GetCharacter(id string) (*[]models.Character, error) {
// 	return f.FerretRepository.GetCharacter(id)
// }
