package services

import (
	"table_top/internal/repositories"
)

type Services struct {
	CharacterService       CharacterService
	LocationService        LocationService
	EnemyService           EnemyService
	EnemyMoveService       EnemyMoveService
	WeaponService          WeaponService
	WeaponComboService     WeaponComboService
	ComboCardService       ComboCardService
	DamageComboCardService DamageComboCardService
	SpecialEffectService   SpecialEffectService
}

func InitServices(r repositories.Repositories) Services {
	return Services{
		CharacterService:       NewCharacterService(r.CharacterRepository),
		LocationService:        NewLocationService(r.LocationRepository),
		EnemyService:           NewEnemyService(r.EnemyRepository),
		EnemyMoveService:       NewEnemyMoveService(r.EnemyMoveRepository),
		WeaponService:          NewWeaponService(r.WeaponRepository),
		WeaponComboService:     NewWeaponComboService(r.WeaponComboRepository),
		ComboCardService:       NewComboCardService(r.ComboCardRepository),
		DamageComboCardService: NewDamageComboCardService(r.DamageComboCardRepository),
		SpecialEffectService:   NewSpecialEffectService(r.SpecialEffectRepository),
	}
}
