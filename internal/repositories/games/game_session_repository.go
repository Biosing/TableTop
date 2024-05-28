package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	models "table_top/internal/models/games"
)

type GameSessionRepository interface {
	Create(ctx context.Context, enemy *models.GameSession) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.GameSession, error)
}

type gameSessionRepository struct {
	db *gorm.DB
}

func NewGameSessionRepository(db *gorm.DB) GameSessionRepository {
	return &gameSessionRepository{db: db}
}

func (r *gameSessionRepository) Create(ctx context.Context, enemy *models.GameSession) error {
	return r.db.WithContext(ctx).Create(enemy).Error
}

func (r *gameSessionRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.GameSession, error) {
	var gameSession models.GameSession
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&gameSession).Error
	if err != nil {
		return nil, err
	}
	return &gameSession, nil
}
