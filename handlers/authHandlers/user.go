package authHandlers

import (
	"gin/db"
	"gin/sturct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsersHandler(c *gin.Context) {
	var users []sturct.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}
