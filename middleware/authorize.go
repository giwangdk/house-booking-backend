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

			fmt.Println("hihhiihi",token)
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
	if len(str) < 2 {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	fmt.Println(strings.Trim(str[1], ""))
	decodedToken := strings.Trim(str[1], " ")
	fmt.Println(decodedToken)

	t, err := validateToken(decodedToken)
	fmt.Println("validate", t, err)
	if err != nil || !t.Valid {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	claims, ok := t.Claims.(jwt.MapClaims)

	fmt.Println("hiyayya", ok)
	fmt.Println("hiyayya", claims)
	
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
