package services

import (
	"context"

	"github.com/google/uuid"

	"table_top/internal/dtos/requests/weapons"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type WeaponService interface {
	CreateWeapon(ctx context.Context, req *weapons.CreateRequest) (*models.Weapon, error)
	GetWeaponByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error)
	UpdateWeapon(ctx context.Context, id uuid.UUID, req *weapons.UpdateRequest) (*models.Weapon, error)
	DeleteWeapon(ctx context.Context, id uuid.UUID) error
	ListWeapons(ctx context.Context) ([]*models.Weapon, error)
}

type weaponService struct {
	repo repositories.WeaponRepository
}

func NewWeaponService(repo repositories.WeaponRepository) WeaponService {
	return &weaponService{repo: repo}
}

func (s *weaponService) CreateWeapon(ctx context.Context, req *weapons.CreateRequest) (*models.Weapon, error) {
	weapon := &models.Weapon{
		CharacterID: req.CharacterID,
		Name:        req.Name,
		CountCards:  req.CountCards,
		Defense:     req.Defense,
	}

	for _, combo := range req.WeaponCombo {
		weaponCombo := models.WeaponCombo{
			ComboType: combo.ComboType,
			Count:     combo.Count,
			Order:     combo.Order,
		}
		weapon.WeaponCombo = append(weapon.WeaponCombo, weaponCombo)
	}

	if err := s.repo.Create(ctx, weapon); err != nil {
		return nil, err
	}

	return weapon, nil
}

func (s *weaponService) GetWeaponByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *weaponService) UpdateWeapon(ctx context.Context, id uuid.UUID, req *weapons.UpdateRequest) (*models.Weapon, error) {
	weapon, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	weapon.CharacterID = req.CharacterID
	weapon.Name = req.Name
	weapon.CountCards = req.CountCards
	weapon.Defense = req.Defense
	if err := s.repo.Update(ctx, weapon); err != nil {
		return nil, err
	}
	return weapon, nil
}

func (s *weaponService) DeleteWeapon(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *weaponService) ListWeapons(ctx context.Context) ([]*models.Weapon, error) {
	return s.repo.List(ctx)
}
