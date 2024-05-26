package services

import (
	"table_top/internal/repositories"
)

type Services struct {
	CharacterService CharacterService
	LocationService  LocationService
	EnemyService     EnemyService
	EnemyMoveService EnemyMoveService
}

func InitServices(r repositories.Repositories) Services {
	return Services{
		CharacterService: NewCharacterService(r.CharacterRepository),
		LocationService:  NewLocationService(r.LocationRepository),
		EnemyService:     NewEnemyService(r.EnemyRepository),
		EnemyMoveService: NewEnemyMoveService(r.EnemyMoveRepository),
	}
}
