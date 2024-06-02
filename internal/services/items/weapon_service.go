package services

import (
	"context"

	"github.com/google/uuid"

	dtos "table_top/internal/dtos/weapons"
	models "table_top/internal/models/items"
	repositories "table_top/internal/repositories/items"
)

type WeaponService interface {
	CreateWeapon(ctx context.Context, req *dtos.CreateRequest) (*models.Weapon, error)
	GetWeaponByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error)
	UpdateWeapon(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.Weapon, error)
	DeleteWeapon(ctx context.Context, id uuid.UUID) error
	ListWeapons(ctx context.Context) ([]*models.Weapon, error)
}

type weaponService struct {
	repo repositories.WeaponRepository
}

func NewWeaponService(repo repositories.WeaponRepository) WeaponService {
	return &weaponService{repo: repo}
}

func (s *weaponService) CreateWeapon(ctx context.Context, req *dtos.CreateRequest) (*models.Weapon, error) {
	weapon := &models.Weapon{
		CharacterID: req.CharacterID,
		Name:        req.Name,
		CountCards:  req.CountCards,
		Defense:     req.Defense,
	}

	for _, combo := range req.WeaponCombo {
		weaponCombo := models.WeaponCombo{
			Type:  combo.Type,
			Count: combo.Count,
			Order: combo.Order,
		}
		weapon.WeaponCombos = append(weapon.WeaponCombos, weaponCombo)
	}

	if err := s.repo.Create(ctx, weapon); err != nil {
		return nil, err
	}

	return weapon, nil
}

func (s *weaponService) GetWeaponByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *weaponService) UpdateWeapon(ctx context.Context, id uuid.UUID, req *dtos.UpdateRequest) (*models.Weapon, error) {
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
