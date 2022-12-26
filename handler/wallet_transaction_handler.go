package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) GetWalletTransactions(c *gin.Context) {

// 	s := c.Query("s")
// 	sortBy := c.Query("sortBy")
// 	sort := c.Query("sort")
// 	limit := c.Query("limit")
// 	page := c.Query("page")

// 	if s == "" && sortBy == "" && sort == "" && limit == "" && page == "" {
// 		userCtx, ok := c.Get("user")
// 		if !ok {
// 			err := httperror.UnauthorizedError()
// 			c.AbortWithStatusJSON(err.StatusCode, err)
// 			return
// 		}
// 		walletID := userCtx.(dto.UserJWT).WalletID
// 		fmt.Println("user", userCtx)
// 		Wallettransactions, err := h.WallettransactionUsecase.GetWalletTransactionsUser(walletID)
// 		if err := c.ShouldBindJSON(&topUpRequest); err != nil {
// 			if appErr, isAppError := err.(httperror.AppError); isAppError {
// 				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
// 				return
// 			}
// 			serverErr := httperror.InternalServerError(err.Error())
// 			c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"statusCode": http.StatusOK,
// 			"data":       Wallettransactions,
// 			"message":    "WalletTransactions fetched successfully",
// 		})
// 	}

// 	limitInt, _ := strconv.Atoi(limit)
// 	if limit == "" {
// 		limitInt = 0
// 	}
// 	pageInt, _ := strconv.Atoi(page)
// 	if page == "" {
// 		pageInt = 0
// 	}

// 	total, Wallettransactions, err := h.WallettransactionUsecase.GetWalletTransactions(s, sortBy, sort, limitInt, pageInt)

// 	if err := c.ShouldBindJSON(&topUpRequest); err != nil {
// 		if appErr, isAppError := err.(httperror.AppError); isAppError {
// 			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
// 			return
// 		}
// 		serverErr := httperror.InternalServerError(err.Error())
// 		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"statusCode": http.StatusOK,
// 		"data":       Wallettransactions,
// 		"message":    "WalletTransactions fetched successfully",
// 		"total":      total,
// 		"page":       pageInt,
// 		"limit":      limitInt,
// 	})
// }

func (h *Handler) TopUp(c *gin.Context) {
	var topUpRequest = new(dto.TopUpRequest)
	userCtx, ok := c.Get("user")
	if !ok {
		httperror.UnauthorizedError()
	}

	topUpRequest.Sender = userCtx.(dto.UserJWT).ID
	topUpRequest.Recipient = userCtx.(dto.UserJWT).ID
	

	if err := c.ShouldBindJSON(&topUpRequest); err != nil {
		if appErr, isAppError := err.(httperror.AppError); isAppError {
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		serverErr := httperror.InternalServerError(err.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
	}

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
		"data":       topUpResponse,
	})
}

