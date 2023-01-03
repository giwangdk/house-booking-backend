package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TopUp(c *gin.Context) {
	var topUpRequest = new(dto.TopUpRequest)
	if err := c.ShouldBindJSON(&topUpRequest); err != nil {
		if appErr, isAppError := err.(httperror.AppError); isAppError {
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		serverErr := httperror.InternalServerError(err.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
	}
	userCtx, ok := c.Get("user")
	if !ok {
		httperror.UnauthorizedError()
	}

	topUpRequest.Sender = userCtx.(dto.UserJWT).ID
	topUpRequest.Recipient = userCtx.(dto.UserJWT).ID

	topUpResponse, err := h.walletTransaction.TopUp(*topUpRequest)
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
		"data":        topUpResponse,
	})
}
