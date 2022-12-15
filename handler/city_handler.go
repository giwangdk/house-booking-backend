package handler

import (
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.cityUsecase.GetCities()
	if err != nil {
		if appErr, isAppError := err.(httperror.AppError); isAppError {
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		serverErr := httperror.InternalServerError(err.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"data":        cities,
	})
}
