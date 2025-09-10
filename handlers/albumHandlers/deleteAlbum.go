package albumHandlers

import (
	"gin/db"
	"gin/sturct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	result := db.DB.Delete(&sturct.Album{}, "id = ?", id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "альбом не найден"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "альбом удалён"})
}
