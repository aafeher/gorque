package main

import (
	"github.com/aafeher/gorque/handlers"
	"github.com/aafeher/gorque/middlewares"
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	auth := r.Group("/api/auth")
	auth.Use(middlewares.CORS())
	auth.Use(middlewares.RateLimitMiddleware(5, 5*time.Minute, 5*time.Minute))
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	api := r.Group("/api")
	api.Use(middlewares.CORS())
	api.Use(middlewares.JWTAuthMiddleware)
	api.GET("/profile", handlers.GetProfile)
	api.PUT("/profile/name", handlers.UpdateProfileName)
	api.PUT("/profile/password", handlers.UpdateProfilePassword)
	api.POST("/refresh", handlers.RefreshToken)

	api.GET("/configuration", handlers.GetConfiguration)

	api.GET("/device", handlers.GetDeviceList)
	api.GET("/session", handlers.GetSessionList)
	api.GET("/data", handlers.GetData)

	r.GET("/upload", handlers.Upload)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
