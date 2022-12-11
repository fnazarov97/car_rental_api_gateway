package handlers

import (
	"net/http"
	"strconv"

	"car_rental/genprotos/car"
	"car_rental/models"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Summary     Create car
// @Description create a new car
// @Tags        cars
// @Accept      json
// @Produce     json
// @Param       car           body     models.CreateCarModel true  "car body"
// @Param       Authorization header   string                false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.PackedCarModel}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/car [post]
func (h Handler) CreateCar(c *gin.Context) {
	var body models.CreateCarModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	obj, err := h.grpcClients.Car.CreateCar(c.Request.Context(), &car.CreateCarRequest{
		Model:   body.Model,
		Color:   body.Color,
		Year:    body.Year,
		Mileage: body.Mileage,
		BrandId: body.BrandId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(c.Request.Context(), &car.GetCarByIDRequest{
		Id: obj.BrandId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "getcarErr",
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Car | Created",
		Data:    car,
	})
}

// GetCarByID godoc
// @Summary     get car by id
// @Description get an car by id
// @Tags        cars
// @Accept      json
// @Param       id            path   string true  "Car ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedCarModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/car/{id} [get]
func (h Handler) GetCarByID(c *gin.Context) {
	idStr := c.Param("id")

	car, err := h.grpcClients.Car.GetCarByID(c.Request.Context(), &car.GetCarByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "getcarErr",
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Car | Created",
		Data:    car,
	})
}

// GetCarList godoc
// @Summary     List cars
// @Description get cars
// @Tags        cars
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Car}
// @Router      /v1/car [get]
func (h Handler) GetCarList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Conf.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Conf.DefaultLimit)
	searchStr := c.DefaultQuery("search", "")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "offset error",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "limit error",
		})
		return
	}
	carList, err := h.grpcClients.Car.GetCarList(c.Request.Context(), &car.GetCarListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    carList,
	})
}

// UpdateCar godoc
// @Summary     update car
// @Description update a new car
// @Tags        cars
// @Accept      json
// @Produce     json
// @Param       car           body     models.UpdateCarModel true  "car body"
// @Param       Authorization header   string                false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Car}
// @Response    400           {object} models.JSONErrorResponse
// @Router      /v1/car [put]
func (h Handler) UpdateCar(c *gin.Context) {
	var body models.UpdateCarModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	updated, err := h.grpcClients.Car.UpdateCar(c.Request.Context(), &car.UpdateCarRequest{
		Id:      body.CarId,
		Model:   body.Model,
		Color:   body.Color,
		Year:    body.Year,
		Mileage: body.Mileage,
		BrandId: body.BrandId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "Car | Update",
	})
}

// DeleteCar godoc
// @Summary     delete car by id
// @Description delete an car by id
// @Tags        cars
// @Accept      json
// @Param       id            path   string true  "Car ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Car}
// @Failure     404 {object} error
// @Router      /v1/car/{id} [delete]
func (h Handler) DeleteCar(c *gin.Context) {
	idStr := c.Param("id")
	car, err := h.grpcClients.Car.DeleteCar(c.Request.Context(), &car.DeleteCarRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Car | Delete | NOT FOUND",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Car deleted",
		"data":    car,
	})
}
