package services

import (
	"context"

	"github.com/google/uuid"

	"table_top/internal/dtos/requests/enemies"
	"table_top/internal/models"
	"table_top/internal/repositories"
)

type EnemyService interface {
	CreateEnemy(ctx context.Context, req *enemies.CreateRequest) (*models.Enemy, error)
	UpdateEnemy(ctx context.Context, id uuid.UUID, req *enemies.UpdateRequest) (*models.Enemy, error)
	GetEnemyByID(ctx context.Context, id uuid.UUID) (*models.Enemy, error)
	DeleteEnemy(ctx context.Context, id uuid.UUID) error
	ListEnemies(ctx context.Context) ([]*models.Enemy, error)
}

type enemyService struct {
	repo repositories.EnemyRepository
}

func NewEnemyService(repo repositories.EnemyRepository) EnemyService {
	return &enemyService{repo: repo}
}

func (s *enemyService) CreateEnemy(ctx context.Context, req *enemies.CreateRequest) (*models.Enemy, error) {
	enemy := &models.Enemy{
		Name:         req.Name,
		Level:        req.Level,
		Class:        req.Class,
		Description:  req.Description,
		Hp:           req.Hp,
		Experience:   req.Experience,
		QuantityDeck: req.QuantityDeck,
		Defense:      req.Defense,
	}

	for _, move := range req.EnemyMoves {
		enemyMove := models.EnemyMove{
			RangeFrom:   move.RangeFrom,
			RangeTo:     move.RangeTo,
			Description: move.Description,
		}
		enemy.EnemyMoves = append(enemy.EnemyMoves, enemyMove)
	}

	if err := s.repo.Create(ctx, enemy); err != nil {
		return nil, err
	}

	return enemy, nil
}

func (s *enemyService) UpdateEnemy(ctx context.Context, id uuid.UUID, req *enemies.UpdateRequest) (*models.Enemy, error) {
	enemy, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	enemy.Name = req.Name
	enemy.Level = req.Level
	enemy.Class = req.Class
	enemy.Description = req.Description
	enemy.Hp = req.Hp
	enemy.Experience = req.Experience
	enemy.QuantityDeck = req.QuantityDeck
	enemy.Defense = req.Defense

	if err := s.repo.Update(ctx, enemy); err != nil {
		return nil, err
	}

	return enemy, nil
}

func (s *enemyService) GetEnemyByID(ctx context.Context, id uuid.UUID) (*models.Enemy, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *enemyService) DeleteEnemy(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *enemyService) ListEnemies(ctx context.Context) ([]*models.Enemy, error) {
	return s.repo.List(ctx)
}
