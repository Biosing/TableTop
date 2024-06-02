package services

import (
	"context"

	"github.com/google/uuid"

	dtos "table_top/internal/dtos/special_effects"
	models "table_top/internal/models/combo_cards"
	repositories "table_top/internal/repositories/combo_cards"
)

type SpecialEffectService interface {
	CreateSpecialEffect(ctx context.Context, req *dtos.CreateRequest) (*models.SpecialEffect, error)
	DeleteSpecialEffect(ctx context.Context, id uuid.UUID) error
	GetSpecialEffectByID(ctx context.Context, id uuid.UUID) (*models.SpecialEffect, error)
	ListSpecialEffects(ctx context.Context) ([]*models.SpecialEffect, error)
	UpdateSpecialEffect(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.SpecialEffect, error)
}

type specialEffectService struct {
	repo repositories.SpecialEffectRepository
}

func NewSpecialEffectService(repo repositories.SpecialEffectRepository) SpecialEffectService {
	return &specialEffectService{repo: repo}
}

func (s *specialEffectService) CreateSpecialEffect(ctx context.Context, req *dtos.CreateRequest) (*models.SpecialEffect, error) {
	specialEffect := &models.SpecialEffect{
		ComboCardID: req.ComboCardID,
		Type:        req.Type,
		DamageType:  req.DamageType,
		Damage:      req.Damage,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, specialEffect); err != nil {
		return nil, err
	}

	return specialEffect, nil
}

func (s *specialEffectService) DeleteSpecialEffect(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *specialEffectService) GetSpecialEffectByID(ctx context.Context, id uuid.UUID) (*models.SpecialEffect, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *specialEffectService) ListSpecialEffects(ctx context.Context) ([]*models.SpecialEffect, error) {
	return s.repo.List(ctx)
}

func (s *specialEffectService) UpdateSpecialEffect(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.SpecialEffect, error) {
	specialEffect, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	specialEffect.ComboCardID = req.ComboCardID
	specialEffect.Type = req.Type
	specialEffect.DamageType = req.DamageType
	specialEffect.Damage = req.Damage
	specialEffect.Description = req.Description

	if err := s.repo.Update(ctx, specialEffect); err != nil {
		return nil, err
	}

	return specialEffect, nil
}
