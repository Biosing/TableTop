package routers

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "table_top/docs" // Импортируйте сгенерированные Swagger-документы
	"table_top/internal/configs"
)

// InitRouter инициализирует маршрутизатор Gin
func InitRouter(ctx context.Context, cfg *configs.Config) *gin.Engine {
	router := gin.Default()

	// Инициализация сессий
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	router.Use(sessions.Sessions(cfg.SessionName, store))

	// Добавление контекста к каждому запросу
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitRoutes(router)

	return router
}
