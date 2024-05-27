package routers

import (
	"github.com/gin-gonic/gin"

	"table_top/internal/handlers"
)

func InitRoutes(router *gin.Engine) {
	//Маршруты для Characters
	router.POST("/characters", handlers.GetCharacterHandler().CreateCharacter)
	router.DELETE("/characters/:id", handlers.GetCharacterHandler().DeleteCharacter)
	router.GET("/characters/:id", handlers.GetCharacterHandler().GetCharacter)
	router.GET("/characters", handlers.GetCharacterHandler().ListCharacters)
	router.PUT("/characters", handlers.GetCharacterHandler().UpdateCharacter)

	// Маршруты для Locations
	router.POST("/locations", handlers.GetLocationHandler().CreateLocation)
	router.DELETE("/locations/:id", handlers.GetLocationHandler().DeleteLocation)
	router.GET("/locations/:id", handlers.GetLocationHandler().GetLocation)
	router.GET("/locations", handlers.GetLocationHandler().ListLocations)
	router.PUT("/locations", handlers.GetLocationHandler().UpdateLocation)

	//Маршруты для Enemies
	router.POST("/enemies", handlers.GetEnemyHandler().CreateEnemy)
	router.PUT("/enemies/:id", handlers.GetEnemyHandler().UpdateEnemy)
	router.GET("/enemies/:id", handlers.GetEnemyHandler().GetEnemy)
	router.DELETE("/enemies/:id", handlers.GetEnemyHandler().DeleteEnemy)
	router.GET("/enemies", handlers.GetEnemyHandler().ListEnemies)

	//Маршруты для EnemyMoves
	router.POST("/enemy_moves/:enemyId", handlers.GetEnemyMoveHandler().CreateEnemyMove)
	router.PUT("/enemy_moves/:id", handlers.GetEnemyMoveHandler().UpdateEnemyMove)
	router.GET("/enemy_moves/:id", handlers.GetEnemyMoveHandler().GetEnemyMove)
	router.DELETE("/enemy_moves/:id", handlers.GetEnemyMoveHandler().DeleteEnemyMove)
	router.GET("/enemy_moves", handlers.GetEnemyMoveHandler().ListEnemyMoves)

	//Маршруты для Weapon
	router.POST("/weapons", handlers.GetWeaponHandler().CreateWeapon)
	router.GET("/weapons/:id", handlers.GetWeaponHandler().GetWeapon)
	router.PUT("/weapons/:id", handlers.GetWeaponHandler().UpdateWeapon)
	router.DELETE("/weapons/:id", handlers.GetWeaponHandler().DeleteWeapon)
	router.GET("/weapons", handlers.GetWeaponHandler().ListWeapons)

	//Маршруты для Weapon Combo
	router.POST("/weapon_combos", handlers.GetWeaponComboHandler().CreateWeaponCombo)
	router.GET("/weapon_combos/:id", handlers.GetWeaponComboHandler().GetWeaponCombo)
	router.PUT("/weapon_combos/:id", handlers.GetWeaponComboHandler().UpdateWeaponCombo)
	router.DELETE("/weapon_combos/:id", handlers.GetWeaponComboHandler().DeleteWeaponCombo)
	router.GET("/weapon_combos", handlers.GetWeaponComboHandler().ListWeaponCombos)
}
