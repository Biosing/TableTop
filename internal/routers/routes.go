package routers

import (
	"github.com/gin-gonic/gin"

	"table_top/internal/handlers"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/characters", handlers.GetCharacterHandler().CreateCharacter)
	router.DELETE("/characters/:id", handlers.GetCharacterHandler().DeleteCharacter)
	router.GET("/characters/:id", handlers.GetCharacterHandler().GetCharacter)
	router.GET("/characters", handlers.GetCharacterHandler().ListCharacters)
	router.PUT("/characters", handlers.GetCharacterHandler().UpdateCharacter)
}
