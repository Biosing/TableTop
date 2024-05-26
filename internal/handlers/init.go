package handlers

import (
	"table_top/internal/services"
)

var (
	characterHandler *CharacterHandler
	locationHandler  *LocationHandler
	enemyHandler     *EnemyHandler
	enemyMoveHandler *EnemyMoveHandler
)

func InitHandlers(s services.Services) {
	characterHandler = NewCharacterHandler(s.CharacterService)
	locationHandler = NewLocationHandler(s.LocationService)
	enemyHandler = NewEnemyHandler(s.EnemyService)
	enemyMoveHandler = NewEnemyMoveHandler(s.EnemyMoveService)
}

func GetCharacterHandler() *CharacterHandler {
	return characterHandler
}

func GetLocationHandler() *LocationHandler {
	return locationHandler
}

func GetEnemyHandler() *EnemyHandler {
	return enemyHandler
}

func GetEnemyMoveHandler() *EnemyMoveHandler {
	return enemyMoveHandler
}
