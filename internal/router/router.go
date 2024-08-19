package router

import (
	"weather-api/internal/config"
	"weather-api/internal/handlers"
	"weather-api/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes the Gin router with routes and middleware
func SetupRouter(cfg *config.Config) *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Setup middleware
	middleware.SetupMiddleware(router)

	// routes - all the routes to be configured here
	// weather GET route
	router.GET("/weather", handlers.GetWeatherHandler(cfg))

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}