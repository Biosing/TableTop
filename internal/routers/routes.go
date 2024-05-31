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

	//Маршруты для Weapon combo
	router.POST("/weapon_combos", handlers.GetWeaponComboHandler().CreateWeaponCombo)
	router.GET("/weapon_combos/:id", handlers.GetWeaponComboHandler().GetWeaponCombo)
	router.PUT("/weapon_combos/:id", handlers.GetWeaponComboHandler().UpdateWeaponCombo)
	router.DELETE("/weapon_combos/:id", handlers.GetWeaponComboHandler().DeleteWeaponCombo)
	router.GET("/weapon_combos", handlers.GetWeaponComboHandler().ListWeaponCombos)

	//Маршруты для Combo cards
	router.POST("/combo_cards", handlers.GetComboCardHandler().CreateComboCard)
	router.GET("/combo_cards/:id", handlers.GetComboCardHandler().GetComboCard)
	router.PUT("/combo_cards/:id", handlers.GetComboCardHandler().UpdateComboCard)
	router.DELETE("/combo_cards/:id", handlers.GetComboCardHandler().DeleteComboCard)
	router.GET("/combo_cards", handlers.GetComboCardHandler().ListComboCards)

	//Маршруты для Damage combo cards
	router.POST("/damage_combo_cards", handlers.GetDamageComboCardHandler().CreateDamageComboCard)
	router.GET("/damage_combo_cards/:id", handlers.GetDamageComboCardHandler().GetDamageComboCard)
	router.PUT("/damage_combo_cards/:id", handlers.GetDamageComboCardHandler().UpdateDamageComboCard)
	router.DELETE("/damage_combo_cards/:id", handlers.GetDamageComboCardHandler().DeleteDamageComboCard)
	router.GET("/damage_combo_cards", handlers.GetDamageComboCardHandler().ListDamageComboCards)

	//Маршруты для Damage special effects
	router.POST("/special_effects", handlers.GetSpecialEffectHandler().CreateSpecialEffect)
	router.GET("/special_effects/:id", handlers.GetSpecialEffectHandler().GetSpecialEffect)
	router.PUT("/special_effects/:id", handlers.GetSpecialEffectHandler().UpdateSpecialEffect)
	router.DELETE("/special_effects/:id", handlers.GetSpecialEffectHandler().DeleteSpecialEffect)
	router.GET("/special_effects", handlers.GetSpecialEffectHandler().ListSpecialEffects)

	//Маршруты для Game session
	router.POST("/create_game", handlers.GetGameSessionHandler().CreateGameSession)
	router.POST("/start_game", handlers.GetGameSessionHandler().StartGameSession)
	router.GET("/game_session", handlers.GetGameSessionHandler().GetGameSession)
	router.POST("/finish_game", handlers.GetGameSessionHandler().FinishGameSession)
	router.GET("/game_sessions", handlers.GetGameSessionHandler().ListGameSessions)

	//Маршруты для Game session player
	router.POST("/add_player", handlers.GetGameSessionPlayerHandler().AddPlayerToGameSession)
	router.POST("/remove_player", handlers.GetGameSessionPlayerHandler().RemovePlayerFromGameSession)
}
