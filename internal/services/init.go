package services

import (
	"table_top/internal/repositories"
	charactersSrv "table_top/internal/services/characters"
	comboCardsSrv "table_top/internal/services/combo_cards"
	enemiesSrv "table_top/internal/services/enemies"
	gameSessionSrv "table_top/internal/services/games"
	itemsSrv "table_top/internal/services/items"
	locationsSrv "table_top/internal/services/locations"
)

type Services struct {
	CharacterService       charactersSrv.CharacterService
	LocationService        locationsSrv.LocationService
	EnemyService           enemiesSrv.EnemyService
	EnemyMoveService       enemiesSrv.EnemyMoveService
	WeaponService          itemsSrv.WeaponService
	WeaponComboService     itemsSrv.WeaponComboService
	ComboCardService       comboCardsSrv.ComboCardService
	DamageComboCardService comboCardsSrv.DamageComboCardService
	SpecialEffectService   comboCardsSrv.SpecialEffectService
	GameSessionService     gameSessionSrv.GameSessionService
}

func InitServices(r repositories.Repositories) Services {
	return Services{
		CharacterService:       charactersSrv.NewCharacterService(r.CharacterRepository),
		LocationService:        locationsSrv.NewLocationService(r.LocationRepository),
		EnemyService:           enemiesSrv.NewEnemyService(r.EnemyRepository),
		EnemyMoveService:       enemiesSrv.NewEnemyMoveService(r.EnemyMoveRepository),
		WeaponService:          itemsSrv.NewWeaponService(r.WeaponRepository),
		WeaponComboService:     itemsSrv.NewWeaponComboService(r.WeaponComboRepository),
		ComboCardService:       comboCardsSrv.NewComboCardService(r.ComboCardRepository),
		DamageComboCardService: comboCardsSrv.NewDamageComboCardService(r.DamageComboCardRepository),
		SpecialEffectService:   comboCardsSrv.NewSpecialEffectService(r.SpecialEffectRepository),
		GameSessionService:     gameSessionSrv.NewGameSessionService(r.GameSessionRepository),
	}
}
