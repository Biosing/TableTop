package services

import (
	"context"

	"table_top/internal/dtos/requests/locations"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type LocationService interface {
	CreateLocation(ctx context.Context, location *locations.CreateRequest) (*models.Location, error)
	GetLocationByID(ctx context.Context, id string) (*models.Location, error)
	ListLocations(ctx context.Context) ([]*models.Location, error)
	DeleteLocation(ctx context.Context, id string) error
	UpdateLocation(ctx context.Context, location *models.Location) (*models.Location, error)
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) CreateLocation(ctx context.Context, location *locations.CreateRequest) (*models.Location, error) {

	locationRepo := &models.Location{
		Name:          location.Name,
		Level:         location.Level,
		DangerLevel:   location.DangerLevel,
		Description:   location.Description,
		MonstersCount: location.MonstersCount,
	}

	if err := s.repo.Create(ctx, locationRepo); err != nil {
		return nil, err
	}
	return locationRepo, nil
}

func (s *locationService) GetLocationByID(ctx context.Context, id string) (*models.Location, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *locationService) ListLocations(ctx context.Context) ([]*models.Location, error) {
	return s.repo.List(ctx)
}

func (s *locationService) DeleteLocation(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *locationService) UpdateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {
	if err := s.repo.Update(ctx, location); err != nil {
		return nil, err
	}
	return location, nil
}
