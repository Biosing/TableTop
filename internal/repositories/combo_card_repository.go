package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type ComboCardRepository interface {
	Create(ctx context.Context, comboCard *models.ComboCard) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.ComboCard, error)
	List(ctx context.Context) ([]*models.ComboCard, error)
	Update(ctx context.Context, comboCard *models.ComboCard) error
}

type comboCardRepository struct {
	db *gorm.DB
}

func NewComboCardRepository(db *gorm.DB) ComboCardRepository {
	return &comboCardRepository{db: db}
}

func (r *comboCardRepository) Create(ctx context.Context, comboCard *models.ComboCard) error {
	return r.db.WithContext(ctx).Create(comboCard).Error
}

func (r *comboCardRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.ComboCard{}, "id = ?", id).Error
}

func (r *comboCardRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.ComboCard, error) {
	var comboCard models.ComboCard
	if err := r.db.WithContext(ctx).Preload("DamageComboCards").Preload("SpecialEffects").First(&comboCard, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &comboCard, nil
}

func (r *comboCardRepository) List(ctx context.Context) ([]*models.ComboCard, error) {
	var comboCards []*models.ComboCard
	if err := r.db.WithContext(ctx).Preload("DamageComboCards").Preload("SpecialEffects").Find(&comboCards).Error; err != nil {
		return nil, err
	}
	return comboCards, nil
}

func (r *comboCardRepository) Update(ctx context.Context, comboCard *models.ComboCard) error {
	return r.db.WithContext(ctx).Save(comboCard).Error
}
