package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type EnemyMoveRepository interface {
	Create(ctx context.Context, enemyMove *models.EnemyMove) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.EnemyMove, error)
	List(ctx context.Context) ([]*models.EnemyMove, error)
	FindByEnemyID(ctx context.Context, enemyID uuid.UUID) ([]*models.EnemyMove, error)
	Update(ctx context.Context, enemyMove *models.EnemyMove) error
}

type enemyMoveRepository struct {
	db *gorm.DB
}

func NewEnemyMoveRepository(db *gorm.DB) EnemyMoveRepository {
	return &enemyMoveRepository{db: db}
}

func (r *enemyMoveRepository) Create(ctx context.Context, enemyMove *models.EnemyMove) error {
	return r.db.WithContext(ctx).Create(enemyMove).Error
}

func (r *enemyMoveRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.EnemyMove, error) {
	var enemyMove models.EnemyMove
	if err := r.db.WithContext(ctx).First(&enemyMove, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &enemyMove, nil
}

func (r *enemyMoveRepository) Update(ctx context.Context, enemyMove *models.EnemyMove) error {
	return r.db.WithContext(ctx).Save(enemyMove).Error
}

func (r *enemyMoveRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.EnemyMove{}, "id = ?", id).Error
}

func (r *enemyMoveRepository) List(ctx context.Context) ([]*models.EnemyMove, error) {
	var enemyMoves []*models.EnemyMove
	if err := r.db.WithContext(ctx).Find(&enemyMoves).Error; err != nil {
		return nil, err
	}
	return enemyMoves, nil
}

func (r *enemyMoveRepository) FindByEnemyID(ctx context.Context, enemyID uuid.UUID) ([]*models.EnemyMove, error) {
	var enemyMoves []*models.EnemyMove
	if err := r.db.WithContext(ctx).Where("enemy_id = ?", enemyID).Find(&enemyMoves).Error; err != nil {
		return nil, err
	}
	return enemyMoves, nil
}
