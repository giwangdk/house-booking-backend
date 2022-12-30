package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateReservation(c *gin.Context) {
	reservation := new(dto.CreateReservationRequest)
	if err := c.ShouldBindJSON(reservation); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}


	reservationRes, err := h.reservationUsecase.CreateReservationWithUser(*reservation)
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
		"data":        reservationRes,
	})
}

func (h *Handler) GetReservationByBookingCode(c *gin.Context) {
	id := c.Param("id")

	reservation, err := h.reservationUsecase.GetReservationByBookingCode(id)
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
		"data":        reservation,
	})
}

func (h *Handler) GetReservationsByUserId(c *gin.Context){
	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	userId := userCtx.(dto.UserJWT).ID

	reservations, err := h.reservationUsecase.GetReservationByUserId(userId)
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
		"data":        reservations,
	})
}