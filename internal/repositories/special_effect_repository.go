package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type SpecialEffectRepository interface {
	Create(ctx context.Context, specialEffect *models.SpecialEffect) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.SpecialEffect, error)
	List(ctx context.Context) ([]*models.SpecialEffect, error)
	Update(ctx context.Context, specialEffect *models.SpecialEffect) error
}

type specialEffectRepository struct {
	db *gorm.DB
}

func NewSpecialEffectRepository(db *gorm.DB) SpecialEffectRepository {
	return &specialEffectRepository{db: db}
}

func (r *specialEffectRepository) Create(ctx context.Context, specialEffect *models.SpecialEffect) error {
	return r.db.WithContext(ctx).Create(specialEffect).Error
}

func (r *specialEffectRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.SpecialEffect{}, "id = ?", id).Error
}

func (r *specialEffectRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.SpecialEffect, error) {
	var specialEffect models.SpecialEffect
	if err := r.db.WithContext(ctx).First(&specialEffect, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &specialEffect, nil
}

func (r *specialEffectRepository) List(ctx context.Context) ([]*models.SpecialEffect, error) {
	var specialEffects []*models.SpecialEffect
	if err := r.db.WithContext(ctx).Find(&specialEffects).Error; err != nil {
		return nil, err
	}
	return specialEffects, nil
}

func (r *specialEffectRepository) Update(ctx context.Context, specialEffect *models.SpecialEffect) error {
	return r.db.WithContext(ctx).Save(specialEffect).Error
}
