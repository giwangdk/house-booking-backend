package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHouseDetail(c *gin.Context) {
	house := new(dto.CreateHouseDetailRequest)
	if err := c.ShouldBindJSON(house); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	fmt.Println(house)

	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	req := dto.CreateHouseDetailRequest{
		MaxGuest:            house.MaxGuest,
		Bedrooms:            house.Bedrooms,
		Beds:                house.Beds,
		Baths:               house.Baths,
		HouseFacilities:     house.HouseFacilities,
		HouseRules:          house.HouseRules,
		HouseServices:       house.HouseServices,
		BathroomsFacilities: house.BathroomsFacilities,
		HouseID:             houseId,
	}

	houseRes, err := h.houseDetailUsecase.CreateHouseDetail(req)
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

func (h *Handler) UpdateHouseDetail(c *gin.Context) {

	house := new(dto.UpdateHouseDetailRequest)
	if err := c.ShouldBindJSON(house); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	id := c.Param("id")

	fmt.Println(house)

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	req := dto.UpdateHouseDetailRequest{
		MaxGuest:            house.MaxGuest,
		Bedrooms:            house.Bedrooms,
		Beds:                house.Beds,
		Baths:               house.Baths,
		HouseFacilities:     house.HouseFacilities,
		HouseRules:          house.HouseRules,
		HouseServices:       house.HouseServices,
		BathroomsFacilities: house.BathroomsFacilities,
	}

	houseRes, err := h.houseDetailUsecase.UpdateHouseDetail(req, houseId)
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
