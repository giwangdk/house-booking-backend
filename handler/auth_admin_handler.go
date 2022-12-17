package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterAdmin(c *gin.Context) {
	var u = new(dto.RegisterRequest)

	if err := c.ShouldBindJSON(u); err != nil {
		err := httperror.BadRequestError(err.Error(), "ERROR_REGISTER")
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	user, err := h.authAdminUsecase.Register(*u)
	if err != nil {
		if appErr, isAppError := err.(httperror.AppError); isAppError {
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		serverErr := httperror.InternalServerError(err.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusCreated,
		"data":        user,
	})
}

func (h *Handler) LoginAdmin(c *gin.Context) {
	var u = new(dto.LoginRequest)

	if err := c.ShouldBindJSON(u); err != nil {
		err := httperror.BadRequestError(err.Error(), "ERROR_LOGIN")
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	user, err := h.authAdminUsecase.Login(*u)
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
		"data": dto.LoginResponse{
			AccessToken: user.AccessToken,
		},
	})
}
