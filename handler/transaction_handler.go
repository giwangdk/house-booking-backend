package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	transaction := new(dto.CreateTransactionRequest)
	if err := c.ShouldBindJSON(transaction); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}


	transactionRes, err := h.transactionUsecase.CreateTransaction(*transaction)	
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
		"data":        transactionRes,
	})
}

func (h *Handler) CreateTransactionRequestGuest(c *gin.Context) {

	bookingCode:= c.Param("booking_code")
	formHeader, err := c.FormFile("transfer_slip")
	if err != nil {
		httperror.InternalServerError(err.Error())
	}


    formFile, err := formHeader.Open()
	if err != nil {
		httperror.InternalServerError(err.Error())
	}

	
	req := dto.CreateTransactionRequest{
		BookingCode:  bookingCode,
		TransferSlip: formFile,
	}

	transactionRes, err := h.transactionUsecase.CreateTransactionRequestGuest(req)	
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
		"data":        transactionRes,
	})
}