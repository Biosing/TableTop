package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	models "table_top/internal/models/games"
)

type GameSessionPlayerRepository interface {
	AddPlayer(ctx context.Context, player *models.Player) error
	CheckDuplicateNickname(ctx context.Context, gameSessionID uuid.UUID, nickname string) (bool, error)
	CheckDuplicateCharacterID(ctx context.Context, gameSessionID uuid.UUID, characterID uuid.UUID) (bool, error)
	RemovePlayerBySessionIDAndNickname(ctx context.Context, gameSessionID uuid.UUID, nickname string) error
}

type gameSessionPlayerRepository struct {
	db *gorm.DB
}

func NewGameSessionPlayerRepository(db *gorm.DB) GameSessionPlayerRepository {
	return &gameSessionPlayerRepository{db: db}
}

func (r *gameSessionPlayerRepository) AddPlayer(ctx context.Context, player *models.Player) error {
	return r.db.WithContext(ctx).Create(player).Error
}

func (r *gameSessionPlayerRepository) CheckDuplicateNickname(ctx context.Context, gameSessionID uuid.UUID, nickname string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Player{}).
		Where("game_session_id = ? AND nickname = ?", gameSessionID, nickname).
		Count(&count).Error
	return count > 0, err
}

func (r *gameSessionPlayerRepository) CheckDuplicateCharacterID(ctx context.Context, gameSessionID uuid.UUID, characterID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Player{}).
		Where("game_session_id = ? AND character_id = ?", gameSessionID, characterID).
		Count(&count).Error
	return count > 0, err

}

func (r *gameSessionPlayerRepository) RemovePlayerBySessionIDAndNickname(ctx context.Context, gameSessionID uuid.UUID, nickname string) error {
	return r.db.WithContext(ctx).Where("game_session_id = ? AND nickname = ?", gameSessionID, nickname).Delete(&models.Player{}).Error
}
