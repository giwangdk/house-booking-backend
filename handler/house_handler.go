package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHouses(c *gin.Context) {

	sortBy := c.Query("sortBy")
	sort := c.Query("sort")
	searchBy := c.Query("searchBy")
	filterBycity := c.Query("filterByCity")
	page := c.Query("page")
	limit := c.Query("limit")
	checkin := c.Query("checkin_date")
	checkout := c.Query("checkout_date")

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

	filterBycityInt, err := strconv.Atoi(filterBycity)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	if filterBycity == "" {
		filterBycityInt = 0
	}

	houses, err := h.houseUsecase.GetHouses(pageInt, limitInt, sortBy, sort, searchBy, filterBycityInt, checkin, checkout)
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
		"data":        houses,
	})
}

func (h *Handler) CreateHouse(c *gin.Context) {
	house := new(dto.CreateHouseRequest)
	if err := c.ShouldBindJSON(house); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	userId := userCtx.(dto.UserJWT).ID

	req := dto.CreateHouseRequest{
		Name:        house.Name,
		Price:       house.Price,
		Location:    house.Location,
		Description: house.Description,
		CityID:      house.CityID,
		UserID:      userId,
	}

	houseRes, err := h.houseUsecase.CreateHouse(req)
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

func (h *Handler) GetHouseById(c *gin.Context) {
	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	house, err := h.houseUsecase.GetHouseById(houseId)
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
		"data":        house,
	})
}

func (h *Handler) UpdateHouse(c *gin.Context) {

	house := new(dto.UpdateHouseRequest)
	if err := c.ShouldBindJSON(house); err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}
	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	req := dto.UpdateHouseRequest{
		Name:        house.Name,
		Price:       house.Price,
		Location:    house.Location,
		Description: house.Description,
		CityID:      house.CityID,
	}

	houseRes, err := h.houseUsecase.UpdateHouse(req, houseId)
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

func (h *Handler) GetHousesHost(c *gin.Context) {
	fmt.Println("hihih")
	userCtx, ok := c.Get("user")
	if !ok {
		err := httperror.UnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	fmt.Println(userCtx)
	userId := userCtx.(dto.UserJWT).ID
	sortBy := c.Query("sortBy")
	sort := c.Query("sort")
	searchBy := c.Query("searchBy")
	page := c.Query("page")
	limit := c.Query("limit")

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

	houses, err := h.houseUsecase.GetHousesHost(userId, pageInt, limitInt, sortBy, sort, searchBy)
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
		"data":        houses,
	})
}

func (h *Handler) DeleteHouse(c *gin.Context) {
	id := c.Param("id")

	houseId, err := strconv.Atoi(id)
	if err != nil {
		httperror.BadRequestError(err.Error(), "BAD_REQUEST")
	}

	err = h.houseUsecase.DeleteHouse(houseId)
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
		"data":        "Delete house successfully",
	})
}