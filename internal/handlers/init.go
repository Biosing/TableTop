package handlers

import (
	"table_top/internal/services"
)

var (
	characterHandler *CharacterHandler
)

func InitHandlers(s services.Services) {
	characterHandler = NewCharacterHandler(s.CharacterService)
}

func GetCharacterHandler() *CharacterHandler {
	return characterHandler
}
