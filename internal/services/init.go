package services

import (
	"table_top/internal/repositories"
)

type Services struct {
	CharacterService CharacterService
}

func InitServices(r repositories.Repositories) Services {
	return Services{
		CharacterService: NewCharacterService(r.CharacterRepository),
	}
}
