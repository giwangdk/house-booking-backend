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

	

	req := dto.CreateReservationRequest{
		CheckIn: reservation.CheckIn,
		CheckOut: reservation.CheckOut,
		TotalPrice: reservation.TotalPrice,
		Fullname: reservation.Fullname,
		Email: reservation.Email,
		CityID: reservation.CityID,
		HouseID: reservation.HouseID,
	}

	reservationRes, err := h.reservationUsecase.CreateReservationWithUser(req)
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

func (h * Handler) GetReservationByBookingCode(c *gin.Context) {
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