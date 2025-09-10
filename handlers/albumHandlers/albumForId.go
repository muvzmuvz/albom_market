package albumHandlers

import (
	"gin/db"
	"gin/sturct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlbumForId(c *gin.Context) {
	id := c.Param("id")
	var album sturct.Album
	if err := db.DB.First(&album, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "альбом не найден"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
