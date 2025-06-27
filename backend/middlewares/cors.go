package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// CORS returns a middleware handler to configure Cross-Origin Resource Sharing policies for incoming HTTP requests.
func CORS() gin.HandlerFunc {

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	if os.Getenv("ENV") == "development" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{os.Getenv("CORS_ORIGINS")}
	}

	return cors.New(corsConfig)
}
