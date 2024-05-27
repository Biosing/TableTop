package repositories

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"

	models "table_top/internal/models/characters"
)

type CharacterRepository interface {
	Create(ctx context.Context, character *models.Character) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Character, error)
	List(ctx context.Context) ([]*models.Character, error)
	Update(ctx context.Context, character *models.Character) error
}

type characterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
	return &characterRepository{db: db}
}

func (r *characterRepository) Create(ctx context.Context, character *models.Character) error {
	return r.db.WithContext(ctx).Create(character).Error
}

func (r *characterRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Character{}, "id = ?", id).Error
}

func (r *characterRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Character, error) {
	var character models.Character
	if err := r.db.WithContext(ctx).First(&character, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *characterRepository) List(ctx context.Context) ([]*models.Character, error) {
	var characters []*models.Character
	if err := r.db.WithContext(ctx).Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (r *characterRepository) Update(ctx context.Context, character *models.Character) error {
	return r.db.WithContext(ctx).Save(character).Error
}
