package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWalletByUserID(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	userId := userCtx.(dto.UserJWT).ID

	wallet, err := h.walletUsecase.GetWalletByUserID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"data":        wallet,
	})
}