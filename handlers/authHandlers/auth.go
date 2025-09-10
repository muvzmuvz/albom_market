package authHandlers

import (
	"gin/db"
	"gin/errors"
	"gin/sturct"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegUser(c *gin.Context) {
	var newUser sturct.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = utils.GeneratedID()

	newUser.Role = "user"

	validUser, err := errors.ValidateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при хэшировании пароля"})
		return
	}
	validUser.Password = string(hashedPassword)

	if err := db.DB.Create(&validUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       validUser.ID,
		"username": validUser.Username,
		"role":     validUser.Role,
	})
}
func Login(c *gin.Context) {
	var loginReq sturct.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user sturct.User
	if err := db.DB.First(&user, "username = ?", loginReq.Username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный логин или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный логин или пароль"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось сгенерировать токен"})
		return
	}

	c.SetCookie(
		"jwt",
		token,
		3600*24*3,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "логин успешен"})
}
