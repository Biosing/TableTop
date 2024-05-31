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
	List(ctx context.Context, name string, limit int, offset int) ([]*models.GameSession, int, error)
	Update(ctx context.Context, gameSession *models.GameSession) error
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

func (r *gameSessionRepository) List(ctx context.Context, name string, limit int, offset int) ([]*models.GameSession, int, error) {
	var gameSessions []*models.GameSession
	var totalCount int64

	query := r.db.WithContext(ctx)

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Model(&models.GameSession{}).Count(&totalCount)

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&gameSessions).Error
	if err != nil {
		return nil, 0, err
	}
	return gameSessions, int(totalCount), nil
}

func (r *gameSessionRepository) Update(ctx context.Context, gameSession *models.GameSession) error {
	return r.db.WithContext(ctx).Model(gameSession).Where("id = ?", gameSession.ID).Updates(gameSession).Error
}
