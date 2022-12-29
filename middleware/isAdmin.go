package middleware

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	user, _ := c.Get("user")
	if user.(dto.UserJWT).Role != "admin" {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	fmt.Println("role", user.(dto.UserJWT).Role )
	c.Next()
}
