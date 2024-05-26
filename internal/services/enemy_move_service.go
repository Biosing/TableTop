package services

import (
	"context"

	"github.com/google/uuid"

	enemymoves "table_top/internal/dtos/requests/enemy_moves"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type EnemyMoveService interface {
	CreateEnemyMove(ctx context.Context, enemyId uuid.UUID, enemyMove *enemymoves.CreateRequest) (*models.EnemyMove, error)
	DeleteEnemyMove(ctx context.Context, id uuid.UUID) error
	GetEnemyMoveByID(ctx context.Context, id uuid.UUID) (*models.EnemyMove, error)
	ListEnemyMoves(ctx context.Context) ([]*models.EnemyMove, error)
	UpdateEnemyMove(ctx context.Context, id uuid.UUID, enemyMove *enemymoves.UpdateRequest) (*models.EnemyMove, error)
}

type enemyMoveService struct {
	repo repositories.EnemyMoveRepository
}

func NewEnemyMoveService(repo repositories.EnemyMoveRepository) EnemyMoveService {
	return &enemyMoveService{repo: repo}
}

func (s *enemyMoveService) CreateEnemyMove(ctx context.Context, enemyId uuid.UUID, req *enemymoves.CreateRequest) (*models.EnemyMove, error) {
	enemyMove := &models.EnemyMove{
		EnemyID:     enemyId,
		RangeFrom:   req.RangeFrom,
		RangeTo:     req.RangeTo,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, enemyMove); err != nil {
		return nil, err
	}

	return enemyMove, nil
}

func (s *enemyMoveService) DeleteEnemyMove(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *enemyMoveService) GetEnemyMoveByID(ctx context.Context, id uuid.UUID) (*models.EnemyMove, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *enemyMoveService) ListEnemyMoves(ctx context.Context) ([]*models.EnemyMove, error) {
	return s.repo.List(ctx)
}

func (s *enemyMoveService) UpdateEnemyMove(ctx context.Context, id uuid.UUID, req *enemymoves.UpdateRequest) (*models.EnemyMove, error) {
	enemyMove, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	enemyMove.RangeFrom = req.RangeFrom
	enemyMove.RangeTo = req.RangeTo
	enemyMove.Description = req.Description

	if err := s.repo.Update(ctx, enemyMove); err != nil {
		return nil, err
	}

	return enemyMove, nil
}
