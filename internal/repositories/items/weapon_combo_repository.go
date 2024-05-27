package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	models "table_top/internal/models/items"
)

type WeaponComboRepository interface {
	Create(ctx context.Context, weaponCombo *models.WeaponCombo) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.WeaponCombo, error)
	Update(ctx context.Context, weaponCombo *models.WeaponCombo) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*models.WeaponCombo, error)
}

type weaponComboRepository struct {
	db *gorm.DB
}

func NewWeaponComboRepository(db *gorm.DB) WeaponComboRepository {
	return &weaponComboRepository{db: db}
}

func (r *weaponComboRepository) Create(ctx context.Context, weaponCombo *models.WeaponCombo) error {
	return r.db.WithContext(ctx).Create(weaponCombo).Error
}

func (r *weaponComboRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.WeaponCombo, error) {
	var weaponCombo models.WeaponCombo
	if err := r.db.WithContext(ctx).First(&weaponCombo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &weaponCombo, nil
}

func (r *weaponComboRepository) Update(ctx context.Context, weaponCombo *models.WeaponCombo) error {
	return r.db.WithContext(ctx).Save(weaponCombo).Error
}

func (r *weaponComboRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.WeaponCombo{}, "id = ?", id).Error
}

func (r *weaponComboRepository) List(ctx context.Context) ([]*models.WeaponCombo, error) {
	var weaponCombos []*models.WeaponCombo
	if err := r.db.WithContext(ctx).Find(&weaponCombos).Error; err != nil {
		return nil, err
	}
	return weaponCombos, nil
}
