package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "table_top/docs" // Импортируйте сгенерированные Swagger-документы
)

// InitRouter инициализирует маршрутизатор Gin
func InitRouter(ctx context.Context) *gin.Engine {
	router := gin.Default()

	// Добавление контекста к каждому запросу
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitRoutes(router)

	return router
}
