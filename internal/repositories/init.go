package repositories

import (
	"gorm.io/gorm"

	charactersRepo "table_top/internal/repositories/characters"
	comboCardsRepo "table_top/internal/repositories/combo_cards"
	enemiesRepo "table_top/internal/repositories/enemies"
	itemsRepo "table_top/internal/repositories/items"
	locationsRepo "table_top/internal/repositories/locations"
)

type Repositories struct {
	CharacterRepository       charactersRepo.CharacterRepository
	LocationRepository        locationsRepo.LocationRepository
	EnemyRepository           enemiesRepo.EnemyRepository
	EnemyMoveRepository       enemiesRepo.EnemyMoveRepository
	WeaponRepository          itemsRepo.WeaponRepository
	WeaponComboRepository     itemsRepo.WeaponComboRepository
	ComboCardRepository       comboCardsRepo.ComboCardRepository
	DamageComboCardRepository comboCardsRepo.DamageComboCardRepository
	SpecialEffectRepository   comboCardsRepo.SpecialEffectRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	return Repositories{
		CharacterRepository:       charactersRepo.NewCharacterRepository(db),
		LocationRepository:        locationsRepo.NewLocationRepository(db),
		EnemyRepository:           enemiesRepo.NewEnemyRepository(db),
		EnemyMoveRepository:       enemiesRepo.NewEnemyMoveRepository(db),
		WeaponRepository:          itemsRepo.NewWeaponRepository(db),
		WeaponComboRepository:     itemsRepo.NewWeaponComboRepository(db),
		ComboCardRepository:       comboCardsRepo.NewComboCardRepository(db),
		DamageComboCardRepository: comboCardsRepo.NewDamageComboCardRepository(db),
		SpecialEffectRepository:   comboCardsRepo.NewSpecialEffectRepository(db),
	}
}
