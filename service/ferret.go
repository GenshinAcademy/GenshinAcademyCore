package service

import (
	"genshinacademycore/models"
	"genshinacademycore/repository"
)

type FerretServiceInterface interface {
	GetAllCharactersStats() (*[]models.CharacterArtifactStatsProfit, error)
	GetAllCharacters() (*[]models.Character, error)
	GetCharacter(id string) (*[]models.Character, error)
}

type FerretService struct {
	FerretRepository repository.FerretRepositoryInterface
}

func NewFerretService(repoFerret repository.FerretRepositoryInterface) FerretServiceInterface {
	return &FerretService{
		FerretRepository: repoFerret,
	}
}

func (f *FerretService) GetAllCharactersStats() (*[]models.CharacterArtifactStatsProfit, error) {
	return f.FerretRepository.GetAllCharactersStats()
}

func (f *FerretService) GetAllCharacters() (*[]models.Character, error) {
	return f.FerretRepository.GetAllCharacters()
}

func (f *FerretService) GetCharacter(id string) (*[]models.Character, error) {
	return f.FerretRepository.GetCharacter(id)
}
