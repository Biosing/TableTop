package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	CharacterRepository       CharacterRepository
	LocationRepository        LocationRepository
	EnemyRepository           EnemyRepository
	EnemyMoveRepository       EnemyMoveRepository
	WeaponRepository          WeaponRepository
	WeaponComboRepository     WeaponComboRepository
	ComboCardRepository       ComboCardRepository
	DamageComboCardRepository DamageComboCardRepository
	SpecialEffectRepository   SpecialEffectRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	return Repositories{
		CharacterRepository:       NewCharacterRepository(db),
		LocationRepository:        NewLocationRepository(db),
		EnemyRepository:           NewEnemyRepository(db),
		EnemyMoveRepository:       NewEnemyMoveRepository(db),
		WeaponRepository:          NewWeaponRepository(db),
		WeaponComboRepository:     NewWeaponComboRepository(db),
		ComboCardRepository:       NewComboCardRepository(db),
		DamageComboCardRepository: NewDamageComboCardRepository(db),
		SpecialEffectRepository:   NewSpecialEffectRepository(db),
	}
}
