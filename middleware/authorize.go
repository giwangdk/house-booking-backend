package middleware

import (
	"encoding/json"
	"errors"
	"final-project-backend/config"
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {

			//todo create sentinel error or http error
			fmt.Println("hihhiihi",token)
			return nil, errors.New("invalid token signature")
		}
		return []byte("very-secret"), nil
	})
}

func Authorize(c *gin.Context) {
	if config.Config.ENV == "testing" {
		fmt.Println("disable auth")
		return
	}
	authHeader := c.GetHeader("Authorization")

	str := strings.Split(authHeader, "Bearer ")
	fmt.Println(str)
	if len(str) < 2 {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	decodedToken := str[1]

	t, err := validateToken(decodedToken)
	if err != nil || !t.Valid {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	var data dto.UserJWT
	user, _ := json.Marshal(claims["user"])
	err = json.Unmarshal(user, &data)

	if err != nil {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	c.Set("user", data)
}
