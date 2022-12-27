package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePickup(c *gin.Context) {
	pickup := new(dto.CreatePickupRequest)
	if err := c.ShouldBindJSON(pickup); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	req := dto.CreatePickupRequest{
		ReservationID: pickup.ReservationID,
		UserID:        pickup.UserID,
	}

	pickupRes, err := h.pickupUsecase.CreatePickup(req)
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
		"data":        pickupRes,
	})
}