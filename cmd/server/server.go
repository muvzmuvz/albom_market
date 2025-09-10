package server

import (
	"gin/handlers/albumHandlers"
	"gin/handlers/authHandlers"
	"gin/middleware"
	"gin/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func StartServer() error {
	router := gin.Default()

	trustErr := router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(utils.LoggerMiddleware())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // если используете cookie для JWT
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/uploads", "./uploads")

	albumRoutes := router.Group("/albums")
	{
		albumRoutes.GET("", albumHandlers.GetAlbums)
		albumRoutes.GET("/:id", albumHandlers.AlbumForId)
		albumRoutes.GET("/search", albumHandlers.SearchAlbums)
		albumRoutes.POST("", middleware.AuthMiddleware(), albumHandlers.AddAlbum)
		albumRoutes.PUT("/:id", middleware.AuthMiddleware(), albumHandlers.UpdateAlbum)
		albumRoutes.DELETE("/:id", middleware.AuthMiddleware(), albumHandlers.DeleteAlbum)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", middleware.AuthMiddleware(), middleware.RequireRole("admin"),
			authHandlers.GetAllUsersHandler)
	}
	
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authHandlers.RegUser)
		authRoutes.POST("/login", authHandlers.Login)
	}

	err := router.Run(":8080")
	if trustErr != nil {
		return trustErr
	}
	if err != nil {
		panic(err)
	}
	return nil
}
