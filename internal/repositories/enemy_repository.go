package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"table_top/internal/models"
)

type EnemyRepository interface {
	Create(ctx context.Context, enemy *models.Enemy) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Enemy, error)
	List(ctx context.Context) ([]*models.Enemy, error)
	Update(ctx context.Context, enemy *models.Enemy) error
}

type enemyRepository struct {
	db *gorm.DB
}

func NewEnemyRepository(db *gorm.DB) EnemyRepository {
	return &enemyRepository{db: db}
}

func (r *enemyRepository) Create(ctx context.Context, enemy *models.Enemy) error {
	return r.db.WithContext(ctx).Create(enemy).Error
}

func (r *enemyRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Enemy, error) {
	var enemy models.Enemy
	if err := r.db.WithContext(ctx).Preload("EnemyMoves").First(&enemy, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &enemy, nil
}

func (r *enemyRepository) Update(ctx context.Context, enemy *models.Enemy) error {
	return r.db.WithContext(ctx).Save(enemy).Error
}

func (r *enemyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Enemy{}, "id = ?", id).Error
}

func (r *enemyRepository) List(ctx context.Context) ([]*models.Enemy, error) {
	var enemies []*models.Enemy
	if err := r.db.WithContext(ctx).Preload("EnemyMoves").Find(&enemies).Error; err != nil {
		return nil, err
	}
	return enemies, nil
}
