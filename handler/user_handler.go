package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	userId := userCtx.(dto.UserJWT).ID

	user, err := h.userUsecase.GetUser(userId)
	if err != nil {
		if appErr, isAppError := err.(httperror.AppError); isAppError {
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		serverErr := httperror.InternalServerError(err.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
	}

	res := dto.DetailUser{
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
		Role:     user.Role,
		CityID:   user.CityID,
		City:     user.City,
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"data":        res,
	})
}

func (h *Handler) EditUser(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	userId := userCtx.(dto.UserJWT).ID

	var editUserRequest = new(dto.EditUserRequest)
	if err := c.ShouldBindJSON(&editUserRequest); err != nil {
		err := httperror.BadRequestError(err.Error(), "BAD_REQUEST")
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	user, err := h.userUsecase.EditUser(*editUserRequest, userId)
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
		"data":        user,
	})
}