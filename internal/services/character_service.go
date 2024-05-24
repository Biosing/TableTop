package services

import (
	"table_top/internal/dtos/requests/characters"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type CharacterService interface {
	CreateCharacter(character *characters.CreateRequest) (*models.Character, error)
	DeleteCharacter(id string) error
	GetCharacterByID(id string) (*models.Character, error)
	ListCharacters() ([]*models.Character, error)
	UpdateCharacter(character *characters.UpdateRequest) (*models.Character, error)
}

type characterService struct {
	repo repositories.CharacterRepository
}

func NewCharacterService(repo repositories.CharacterRepository) CharacterService {
	return &characterService{repo: repo}
}

func (s *characterService) CreateCharacter(character *characters.CreateRequest) (*models.Character, error) {
	repoCharacter := &models.Character{
		Name:         character.Name,
		Class:        character.Class,
		Race:         character.Race,
		Age:          character.Age,
		Sex:          character.Sex,
		Description:  character.Description,
		BeginningWay: character.BeginningWay,
	}

	if err := s.repo.Create(repoCharacter); err != nil {
		return nil, err
	}

	return repoCharacter, nil
}

func (s *characterService) DeleteCharacter(id string) error {
	return s.repo.Delete(id)
}

func (s *characterService) GetCharacterByID(id string) (*models.Character, error) {
	return s.repo.GetByID(id)
}

func (s *characterService) ListCharacters() ([]*models.Character, error) {
	return s.repo.List()
}

func (s *characterService) UpdateCharacter(request *characters.UpdateRequest) (*models.Character, error) {
	character, err := s.repo.GetByID(request.ID)
	if err != nil {
		return nil, err
	}

	character.Name = request.Name
	character.Class = request.Class
	character.Race = request.Race
	character.Age = request.Age
	character.Sex = request.Sex
	character.Description = request.Description
	character.BeginningWay = request.BeginningWay

	if err := s.repo.Update(character); err != nil {
		return nil, err
	}

	return character, nil
}
