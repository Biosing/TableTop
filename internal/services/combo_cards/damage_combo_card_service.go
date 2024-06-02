package services

import (
	"context"

	"github.com/google/uuid"

	dtos "table_top/internal/dtos/damage_combo_cards"
	models "table_top/internal/models/combo_cards"
	repositories "table_top/internal/repositories/combo_cards"
)

type DamageComboCardService interface {
	CreateDamageComboCard(ctx context.Context, req *dtos.CreateRequest) (*models.DamageComboCard, error)
	DeleteDamageComboCard(ctx context.Context, id uuid.UUID) error
	GetDamageComboCardByID(ctx context.Context, id uuid.UUID) (*models.DamageComboCard, error)
	ListDamageComboCards(ctx context.Context) ([]*models.DamageComboCard, error)
	UpdateDamageComboCard(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.DamageComboCard, error)
}

type damageComboCardService struct {
	repo repositories.DamageComboCardRepository
}

func NewDamageComboCardService(repo repositories.DamageComboCardRepository) DamageComboCardService {
	return &damageComboCardService{repo: repo}
}

func (s *damageComboCardService) CreateDamageComboCard(ctx context.Context, req *dtos.CreateRequest) (*models.DamageComboCard, error) {
	damageComboCard := &models.DamageComboCard{
		ComboCardID: req.ComboCardID,
		RangeFrom:   req.RangeFrom,
		RangeTo:     req.RangeTo,
		Type:        req.Type,
		Damage:      req.Damage,
	}

	if err := s.repo.Create(ctx, damageComboCard); err != nil {
		return nil, err
	}

	return damageComboCard, nil
}

func (s *damageComboCardService) DeleteDamageComboCard(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *damageComboCardService) GetDamageComboCardByID(ctx context.Context, id uuid.UUID) (*models.DamageComboCard, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *damageComboCardService) ListDamageComboCards(ctx context.Context) ([]*models.DamageComboCard, error) {
	return s.repo.List(ctx)
}

func (s *damageComboCardService) UpdateDamageComboCard(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.DamageComboCard, error) {
	damageComboCard, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	damageComboCard.ComboCardID = req.ComboCardID
	damageComboCard.RangeFrom = req.RangeFrom
	damageComboCard.RangeTo = req.RangeTo
	damageComboCard.Type = req.Type
	damageComboCard.Damage = req.Damage

	if err := s.repo.Update(ctx, damageComboCard); err != nil {
		return nil, err
	}

	return damageComboCard, nil
}
