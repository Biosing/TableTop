package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "table_top/docs" // Импортируйте сгенерированные Swagger-документы
)

// InitRouter инициализирует маршрутизатор Gin
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitRoutes(router)

	return router
}
