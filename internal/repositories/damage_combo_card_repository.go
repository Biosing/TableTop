package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type DamageComboCardRepository interface {
	Create(ctx context.Context, damageComboCard *models.DamageComboCard) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.DamageComboCard, error)
	List(ctx context.Context) ([]*models.DamageComboCard, error)
	Update(ctx context.Context, damageComboCard *models.DamageComboCard) error
}

type damageComboCardRepository struct {
	db *gorm.DB
}

func NewDamageComboCardRepository(db *gorm.DB) DamageComboCardRepository {
	return &damageComboCardRepository{db: db}
}

func (r *damageComboCardRepository) Create(ctx context.Context, damageComboCard *models.DamageComboCard) error {
	return r.db.WithContext(ctx).Create(damageComboCard).Error
}

func (r *damageComboCardRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.DamageComboCard{}, "id = ?", id).Error
}

func (r *damageComboCardRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.DamageComboCard, error) {
	var damageComboCard models.DamageComboCard
	if err := r.db.WithContext(ctx).First(&damageComboCard, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &damageComboCard, nil
}

func (r *damageComboCardRepository) List(ctx context.Context) ([]*models.DamageComboCard, error) {
	var damageComboCards []*models.DamageComboCard
	if err := r.db.WithContext(ctx).Find(&damageComboCards).Error; err != nil {
		return nil, err
	}
	return damageComboCards, nil
}

func (r *damageComboCardRepository) Update(ctx context.Context, damageComboCard *models.DamageComboCard) error {
	return r.db.WithContext(ctx).Save(damageComboCard).Error
}
