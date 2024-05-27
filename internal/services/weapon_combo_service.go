package services

import (
	"context"

	"github.com/google/uuid"

	"table_top/internal/dtos/requests/weapon_combos"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type WeaponComboService interface {
	CreateWeaponCombo(ctx context.Context, req *weapon_combos.CreateRequest) (*models.WeaponCombo, error)
	GetWeaponComboByID(ctx context.Context, id uuid.UUID) (*models.WeaponCombo, error)
	UpdateWeaponCombo(ctx context.Context, id uuid.UUID, req *weapon_combos.UpdateRequest) (*models.WeaponCombo, error)
	DeleteWeaponCombo(ctx context.Context, id uuid.UUID) error
	ListWeaponCombos(ctx context.Context) ([]*models.WeaponCombo, error)
}

type weaponComboService struct {
	repo repositories.WeaponComboRepository
}

func NewWeaponComboService(repo repositories.WeaponComboRepository) WeaponComboService {
	return &weaponComboService{repo: repo}
}

func (s *weaponComboService) CreateWeaponCombo(ctx context.Context, req *weapon_combos.CreateRequest) (*models.WeaponCombo, error) {
	weaponCombo := &models.WeaponCombo{
		WeaponID:  req.WeaponID,
		ComboType: req.ComboType,
		Count:     req.Count,
		Order:     req.Order,
	}
	if err := s.repo.Create(ctx, weaponCombo); err != nil {
		return nil, err
	}
	return weaponCombo, nil
}

func (s *weaponComboService) GetWeaponComboByID(ctx context.Context, id uuid.UUID) (*models.WeaponCombo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *weaponComboService) UpdateWeaponCombo(ctx context.Context, id uuid.UUID, req *weapon_combos.UpdateRequest) (*models.WeaponCombo, error) {
	weaponCombo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	weaponCombo.WeaponID = req.WeaponID
	weaponCombo.ComboType = req.ComboType
	weaponCombo.Count = req.Count
	weaponCombo.Order = req.Order
	if err := s.repo.Update(ctx, weaponCombo); err != nil {
		return nil, err
	}
	return weaponCombo, nil
}

func (s *weaponComboService) DeleteWeaponCombo(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *weaponComboService) ListWeaponCombos(ctx context.Context) ([]*models.WeaponCombo, error) {
	return s.repo.List(ctx)
}
