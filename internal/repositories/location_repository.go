package repositories

import (
	"context"

	"table_top/internal/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(ctx context.Context, location *models.Location) error
	GetByID(ctx context.Context, id string) (*models.Location, error)
	List(ctx context.Context) ([]*models.Location, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, location *models.Location) error
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) Create(ctx context.Context, location *models.Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *locationRepository) GetByID(ctx context.Context, id string) (*models.Location, error) {
	var location models.Location
	if err := r.db.WithContext(ctx).First(&location, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *locationRepository) List(ctx context.Context) ([]*models.Location, error) {
	var locations []*models.Location
	if err := r.db.WithContext(ctx).Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *locationRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Location{}, "id = ?", id).Error
}

func (r *locationRepository) Update(ctx context.Context, location *models.Location) error {
	return r.db.WithContext(ctx).Save(location).Error
}
