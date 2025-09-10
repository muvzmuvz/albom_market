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

func AddAlbum(c *gin.Context) {
	var newAlbum sturct.Album

	if c.ContentType() == "application/json" {
		if err := c.BindJSON(&newAlbum); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		newAlbum.Title = strings.TrimSpace(c.PostForm("title"))
		newAlbum.Artist = strings.TrimSpace(c.PostForm("artist"))

		file, err := c.FormFile("image")
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "не удалось получить изображение"})
			return
		}

		filename := fmt.Sprintf("%s_%s", utils.GeneratedID(), file.Filename)
		filePath := "uploads/" + filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить файл"})
			return
		}
		newAlbum.ImagePath = "/" + filePath
	}

	newAlbum.ID = utils.GeneratedID()

	validAlbum, err := errors.ValidateAlbums(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&validAlbum).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, validAlbum)
}
