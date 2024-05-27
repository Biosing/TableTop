package services

import (
	"context"

	"github.com/google/uuid"

	"table_top/internal/dtos/requests/characters"
	models "table_top/internal/models/characters"
	repositories "table_top/internal/repositories/characters"
)

type CharacterService interface {
	CreateCharacter(ctx context.Context, character *characters.CreateRequest) (*models.Character, error)
	DeleteCharacter(ctx context.Context, id uuid.UUID) error
	GetCharacterByID(ctx context.Context, id uuid.UUID) (*models.Character, error)
	ListCharacters(ctx context.Context) ([]*models.Character, error)
	UpdateCharacter(ctx context.Context, id uuid.UUID, character *characters.UpdateRequest) (*models.Character, error)
}

type characterService struct {
	repo repositories.CharacterRepository
}

func NewCharacterService(repo repositories.CharacterRepository) CharacterService {
	return &characterService{repo: repo}
}

func (s *characterService) CreateCharacter(ctx context.Context, character *characters.CreateRequest) (*models.Character, error) {
	repoCharacter := &models.Character{
		Name:         character.Name,
		Class:        character.Class,
		Race:         character.Race,
		Age:          character.Age,
		Sex:          character.Sex,
		Description:  character.Description,
		BeginningWay: character.BeginningWay,
	}

	if err := s.repo.Create(ctx, repoCharacter); err != nil {
		return nil, err
	}

	return repoCharacter, nil
}

func (s *characterService) DeleteCharacter(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *characterService) GetCharacterByID(ctx context.Context, id uuid.UUID) (*models.Character, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *characterService) ListCharacters(ctx context.Context) ([]*models.Character, error) {
	return s.repo.List(ctx)
}

func (s *characterService) UpdateCharacter(ctx context.Context, id uuid.UUID, request *characters.UpdateRequest) (*models.Character, error) {
	character, err := s.repo.GetByID(ctx, id)
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

	if err := s.repo.Update(ctx, character); err != nil {
		return nil, err
	}

	return character, nil
}
