package repositories

import (
	"gorm.io/gorm"

	"table_top/internal/models"
)

type CharacterRepository interface {
	Create(character *models.Character) error
	Delete(id string) error
	GetByID(id string) (*models.Character, error)
	List() ([]*models.Character, error)
	Update(character *models.Character) error
}

type characterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
	return &characterRepository{db: db}
}

func (r *characterRepository) Create(character *models.Character) error {
	return r.db.Create(character).Error
}

func (r *characterRepository) Delete(id string) error {
	return r.db.Delete(&models.Character{}, "id = ?", id).Error
}

func (r *characterRepository) GetByID(id string) (*models.Character, error) {
	var character models.Character
	if err := r.db.First(&character, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *characterRepository) List() ([]*models.Character, error) {
	var characters []*models.Character
	if err := r.db.Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (r *characterRepository) Update(character *models.Character) error {
	return r.db.Save(character).Error
}
