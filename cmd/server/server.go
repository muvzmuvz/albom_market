package server

import (
	"gin/handlers/albumHandlers"
	"gin/handlers/authHandlers"
	"gin/middleware"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	router := gin.Default()
	trustErr := router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.Use(utils.LoggerMiddleware())
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
