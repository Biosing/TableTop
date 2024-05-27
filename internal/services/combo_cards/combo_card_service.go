package services

import (
	"context"

	"table_top/internal/dtos/requests/combo_cards"
	models "table_top/internal/models/combo_cards"
	repositories "table_top/internal/repositories/combo_cards"

	"github.com/google/uuid"
)

type ComboCardService interface {
	CreateComboCard(ctx context.Context, req *combo_cards.CreateRequest) (*models.ComboCard, error)
	DeleteComboCard(ctx context.Context, id uuid.UUID) error
	GetComboCardByID(ctx context.Context, id uuid.UUID) (*models.ComboCard, error)
	ListComboCards(ctx context.Context) ([]*models.ComboCard, error)
	UpdateComboCard(ctx context.Context, id uuid.UUID, req *combo_cards.UpdateRequest) (*models.ComboCard, error)
}

type comboCardService struct {
	repo repositories.ComboCardRepository
}

func NewComboCardService(repo repositories.ComboCardRepository) ComboCardService {
	return &comboCardService{repo: repo}
}

func (s *comboCardService) CreateComboCard(ctx context.Context, req *combo_cards.CreateRequest) (*models.ComboCard, error) {
	comboCard := &models.ComboCard{
		CharacterID:         req.CharacterID,
		Type:                req.Type,
		Name:                req.Name,
		Description:         req.Description,
		TargetEnemyID:       req.TargetEnemyID,
		RequiredNumberCells: req.RequiredNumberCells,
		AddedNumberCells:    req.AddedNumberCells,
	}

	for _, damage := range req.DamageComboCards {
		damageComboCard := models.DamageComboCard{
			RangeFrom: damage.RangeFrom,
			RangeTo:   damage.RangeTo,
			Type:      damage.Type,
			Damage:    damage.Damage,
		}
		comboCard.DamageComboCards = append(comboCard.DamageComboCards, damageComboCard)
	}

	for _, effect := range req.SpecialEffects {
		specialEffect := models.SpecialEffect{
			Type:        effect.Type,
			DamageType:  effect.DamageType,
			Damage:      effect.Damage,
			Description: effect.Description,
		}
		comboCard.SpecialEffects = append(comboCard.SpecialEffects, specialEffect)
	}

	if err := s.repo.Create(ctx, comboCard); err != nil {
		return nil, err
	}

	return comboCard, nil
}

func (s *comboCardService) DeleteComboCard(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *comboCardService) GetComboCardByID(ctx context.Context, id uuid.UUID) (*models.ComboCard, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *comboCardService) ListComboCards(ctx context.Context) ([]*models.ComboCard, error) {
	return s.repo.List(ctx)
}

func (s *comboCardService) UpdateComboCard(ctx context.Context, id uuid.UUID, req *combo_cards.UpdateRequest) (*models.ComboCard, error) {
	comboCard, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	comboCard.CharacterID = req.CharacterID
	comboCard.Type = req.Type
	comboCard.Name = req.Name
	comboCard.Description = req.Description
	comboCard.TargetEnemyID = req.TargetEnemyID
	comboCard.RequiredNumberCells = req.RequiredNumberCells
	comboCard.AddedNumberCells = req.AddedNumberCells

	// Обновление DamageComboCards и SpecialEffects по аналогии с созданием

	if err := s.repo.Update(ctx, comboCard); err != nil {
		return nil, err
	}

	return comboCard, nil
}
