package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHousePhoto(c *gin.Context) {
    formHeader, err := c.FormFile("photo")
	if err != nil {
		httperror.InternalServerError(err.Error())
	}

    formFile, err := formHeader.Open()
	if err != nil {
		httperror.InternalServerError(err.Error())
	}

	house := new(dto.CreateHousePhotoRequest)
	if err := c.ShouldBindJSON(house); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	req := dto.CreateHousePhotoRequest{
		HouseID: houseId,
		Photo:   formFile,
	}

	houseRes, err := h.housePhotoUsecase.CreateHousePhoto(req)
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
		"data":        houseRes,
	})
}

func (h *Handler) DeleteHousePhoto(c *gin.Context) {
	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	err = h.housePhotoUsecase.DeleteHousePhoto(houseId)
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
	})
}
