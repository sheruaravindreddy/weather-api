package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetupMiddleware adds middleware to the Gin router
// Additional layers like access-control or observability could be integrated here
func SetupMiddleware(r *gin.Engine) {
	// Logger middleware
	r.Use(gin.Logger())

	// Recovery middleware
	r.Use(gin.Recovery())
}