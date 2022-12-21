package handler

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHouses(c *gin.Context) {

	sortBy:= c.Query("sortBy")
	sort:= c.Query("sort")
	searchBy:= c.Query("searchBy")
	filterBycity:= c.Query("filterByCity")
	page:= c.Query("page")
	limit:= c.Query("limit")


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

	houses, err := h.houseUsecase.GetHouses(pageInt, limitInt, sortBy, sort, searchBy, filterBycityInt)
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


	req:= dto.CreateHouseRequest{
		Name: house.Name,
		Price: house.Price,
		Location: house.Location,
		Description: house.Description,
		CityID: house.CityID,
		UserID: userId,
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