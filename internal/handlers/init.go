package handlers

import (
	"table_top/internal/services"
)

var (
	characterHandler       *CharacterHandler
	locationHandler        *LocationHandler
	enemyHandler           *EnemyHandler
	enemyMoveHandler       *EnemyMoveHandler
	weaponHandler          *WeaponHandler
	weaponComboHandler     *WeaponComboHandler
	comboCardHandler       *ComboCardHandler
	damageComboCardHandler *DamageComboCardHandler
	specialEffectHandler   *SpecialEffectHandler
	gameSessionHandler     *GameSessionHandler
)

func InitHandlers(s services.Services) {
	characterHandler = NewCharacterHandler(s.CharacterService)
	locationHandler = NewLocationHandler(s.LocationService)
	enemyHandler = NewEnemyHandler(s.EnemyService)
	enemyMoveHandler = NewEnemyMoveHandler(s.EnemyMoveService)
	weaponHandler = NewWeaponHandler(s.WeaponService)
	weaponComboHandler = NewWeaponComboHandler(s.WeaponComboService)
	comboCardHandler = NewComboCardHandler(s.ComboCardService)
	damageComboCardHandler = NewDamageComboCardHandler(s.DamageComboCardService)
	specialEffectHandler = NewSpecialEffectHandler(s.SpecialEffectService)
	gameSessionHandler = NewGameSessionHandler(s.GameSessionService)
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

func GetComboCardHandler() *ComboCardHandler {
	return comboCardHandler
}

func GetDamageComboCardHandler() *DamageComboCardHandler {
	return damageComboCardHandler
}

func GetSpecialEffectHandler() *SpecialEffectHandler {
	return specialEffectHandler
}

func GetGameSessionHandler() *GameSessionHandler { return gameSessionHandler }
