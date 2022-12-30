package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetPickups(c *gin.Context) {

	sortBy := c.Query("sortBy")
	sort := c.Query("sort")
	searchBy := c.Query("searchBy")
	page := c.Query("page")
	limit := c.Query("limit")
	pickupStatusId:= c.Query("pickup_status")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	if limit == "" {
		limitInt = 0
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	if page == "" {
		pageInt = 0
	}

	pickupStatusIdInt, err := strconv.Atoi(pickupStatusId)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	if pickupStatusId == "" {
		pickupStatusIdInt = 0
	}


	pickups, err := h.pickupUsecase.GetPickups(pageInt, limitInt, sortBy, sort, searchBy, pickupStatusIdInt)
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
		"data":        pickups,
	})
}


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

func (h *Handler) UpdateStatusPickup(c *gin.Context) {	

	pickup := new(dto.UpdateStatusPickupRequest)
	if err := c.ShouldBindJSON(pickup); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	id:= c.Param("id")
	pickupId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	pickupRes, err := h.pickupUsecase.UpdateStatusPickup(pickupId,pickup.PickupStatusID)
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