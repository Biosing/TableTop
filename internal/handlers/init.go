package handlers

import (
	"table_top/internal/services"
)

var (
	characterHandler   *CharacterHandler
	locationHandler    *LocationHandler
	enemyHandler       *EnemyHandler
	enemyMoveHandler   *EnemyMoveHandler
	weaponHandler      *WeaponHandler
	weaponComboHandler *WeaponComboHandler
)

func InitHandlers(s services.Services) {
	characterHandler = NewCharacterHandler(s.CharacterService)
	locationHandler = NewLocationHandler(s.LocationService)
	enemyHandler = NewEnemyHandler(s.EnemyService)
	enemyMoveHandler = NewEnemyMoveHandler(s.EnemyMoveService)
	weaponHandler = NewWeaponHandler(s.WeaponService)
	weaponComboHandler = NewWeaponComboHandler(s.WeaponComboService)
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

func GetWeaponHandler() *WeaponHandler {
	return weaponHandler
}

func GetWeaponComboHandler() *WeaponComboHandler {
	return weaponComboHandler
}
