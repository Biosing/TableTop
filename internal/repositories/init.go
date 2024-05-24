package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	CharacterRepository CharacterRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	return Repositories{
		CharacterRepository: NewCharacterRepository(db),
	}
}
