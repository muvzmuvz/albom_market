package albumHandlers

import (
	"gin/db"
	"gin/sturct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SearchAlbums(c *gin.Context) {
	title := strings.TrimSpace(c.Query("title"))
	artist := strings.TrimSpace(c.Query("artist"))

	if title == "" && artist == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "необходимо указать title или artist для поиска"})
		return
	}

	var albums []sturct.Album
	query := db.DB
	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}
	if artist != "" {
		query = query.Where("LOWER(artist) LIKE ?", "%"+strings.ToLower(artist)+"%")
	}
	if err := query.Find(&albums).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}
