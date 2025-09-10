package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func GeneratedID() string {
	return uuid.New().String()
}

//тест middleware

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Print("Новый запрос: ", path, "	", method, " ", time.Since(start))
		c.Next()
	}
}
