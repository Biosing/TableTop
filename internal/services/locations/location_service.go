package services

import (
	"context"

	"github.com/google/uuid"

	"table_top/internal/dtos/requests/locations"
	locations2 "table_top/internal/models/locations"
	locations3 "table_top/internal/repositories/locations"
)

type LocationService interface {
	CreateLocation(ctx context.Context, location *locations.CreateRequest) (*locations2.Location, error)
	GetLocationByID(ctx context.Context, id uuid.UUID) (*locations2.Location, error)
	ListLocations(ctx context.Context) ([]*locations2.Location, error)
	DeleteLocation(ctx context.Context, id uuid.UUID) error
	UpdateLocation(ctx context.Context, id uuid.UUID, location *locations.UpdateRequest) (*locations2.Location, error)
}

type locationService struct {
	repo locations3.LocationRepository
}

func NewLocationService(repo locations3.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) CreateLocation(ctx context.Context, location *locations.CreateRequest) (*locations2.Location, error) {

	locationRepo := &locations2.Location{
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

func (s *locationService) GetLocationByID(ctx context.Context, id uuid.UUID) (*locations2.Location, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *locationService) ListLocations(ctx context.Context) ([]*locations2.Location, error) {
	return s.repo.List(ctx)
}

func (s *locationService) DeleteLocation(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *locationService) UpdateLocation(ctx context.Context, id uuid.UUID, req *locations.UpdateRequest) (*locations2.Location, error) {
	loc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	loc.ID = id
	loc.Name = req.Name
	loc.Level = req.Level
	loc.DangerLevel = req.DangerLevel
	loc.Description = req.Description
	loc.MonstersCount = req.MonstersCount

	if err := s.repo.Update(ctx, loc); err != nil {
		return nil, err
	}

	return loc, nil
}
