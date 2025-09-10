package albumHandlers

import (
	"fmt"
	"gin/db"
	"gin/errors"
	"gin/sturct"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var existingAlbum sturct.Album
	if err := db.DB.First(&existingAlbum, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "альбом не найден"})
		return
	}

	if c.ContentType() == "application/json" {
		var input struct {
			Title  *string `json:"title"`
			Artist *string `json:"artist"`
			Desc   *string `json:"desc"`
			Price  *int    `json:"price"`
		}
		if err := c.BindJSON(&input); err == nil {
			if input.Title != nil {
				existingAlbum.Title = *input.Title
			}
			if input.Artist != nil {
				existingAlbum.Artist = *input.Artist
			}
			if input.Desc != nil {
				existingAlbum.Desc = *input.Desc
			}
			if input.Price != nil {
				existingAlbum.Price = *input.Price
			}
		}
	} else {
		title := strings.TrimSpace(c.PostForm("title"))
		artist := strings.TrimSpace(c.PostForm("artist"))
		if title != "" {
			existingAlbum.Title = title
		}
		if artist != "" {
			existingAlbum.Artist = artist
		}

		file, err := c.FormFile("image")
		if err == nil {
			filename := fmt.Sprintf("%s_%s", utils.GeneratedID(), file.Filename)
			filePath := "uploads/" + filename
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить файл"})
				return
			}
			existingAlbum.ImagePath = "/" + filePath
		}
	}

	if err := errors.ValidateUpdateAlbum(existingAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&existingAlbum).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, existingAlbum)
}
