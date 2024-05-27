package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type WeaponRepository interface {
	Create(ctx context.Context, weapon *models.Weapon) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error)
	Update(ctx context.Context, weapon *models.Weapon) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*models.Weapon, error)
}

type weaponRepository struct {
	db *gorm.DB
}

func NewWeaponRepository(db *gorm.DB) WeaponRepository {
	return &weaponRepository{db: db}
}

func (r *weaponRepository) Create(ctx context.Context, weapon *models.Weapon) error {
	return r.db.WithContext(ctx).Create(weapon).Error
}

func (r *weaponRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Weapon, error) {
	var weapon models.Weapon
	if err := r.db.WithContext(ctx).First(&weapon, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &weapon, nil
}

func (r *weaponRepository) Update(ctx context.Context, weapon *models.Weapon) error {
	return r.db.WithContext(ctx).Save(weapon).Error
}

func (r *weaponRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Weapon{}, "id = ?", id).Error
}

func (r *weaponRepository) List(ctx context.Context) ([]*models.Weapon, error) {
	var weapons []*models.Weapon
	if err := r.db.WithContext(ctx).Find(&weapons).Error; err != nil {
		return nil, err
	}
	return weapons, nil
}
