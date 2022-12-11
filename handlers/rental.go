package handlers

import (
	"car_rental/genprotos/authorization"
	"car_rental/genprotos/car"
	"car_rental/genprotos/rental"
	"car_rental/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRental godoc
// @Summary     Create rental
// @Description create a new rental
// @Tags        rentals
// @Accept      json
// @Produce     json
// @Param       rental        body     models.CreateRentalModel true  "rental body"
// @Param       Authorization header   string                   false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Rental}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/rental [post]
func (h Handler) CreateRental(c *gin.Context) {
	var body models.CreateRentalModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(c.Request.Context(), &car.GetCarByIDRequest{
		Id: body.CarId,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Car not found to create rental",
		})
		return
	}

	customer, err := h.grpcClients.Authorization.GetUserByID(c.Request.Context(), &authorization.GetUserByIDRequest{
		Id: body.CustomerId,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "User is not registired to rental car",
		})
		return
	}

	rental, err := h.grpcClients.Rental.CreateRental(c.Request.Context(), &rental.CreateRentalRequest{
		CarId:      car.CarId,
		CustomerId: customer.Id,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
		Payment:    body.Payment,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Rental | Created",
		Data:    rental,
	})
}

// GetRentalByID godoc
// @Summary     get rental by id
// @Description get an rental by id
// @Tags        rentals
// @Accept      json
// @Param       id            path   string true  "Rental ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedRentalModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/rental/{id} [get]
func (h Handler) GetRentalByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation
	rental, err := h.grpcClients.Rental.GetRentalByID(c.Request.Context(), &rental.GetRentalByIDRequest{
		RentalId: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	car, err := h.grpcClients.Car.GetCarByID(c.Request.Context(), &car.GetCarByIDRequest{
		Id: rental.CarId,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "failed to get car",
		})
		return
	}
	rental.Car.CarId = car.CarId
	rental.Car.Model = car.Model
	rental.Car.Color = car.Color
	rental.Car.Year = car.Year
	rental.Car.Mileage = car.Mileage
	rental.Car.BrandId = car.Brand.BrandId
	rental.Car.CreatedAt = car.CreatedAt
	rental.Car.UpdatedAt = car.UpdatedAt

	customer, err := h.grpcClients.Authorization.GetUserByID(c.Request.Context(), &authorization.GetUserByIDRequest{
		Id: rental.CustomerId,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "failed to get user",
		})
		return
	}
	rental.Customer.Id = customer.Id
	rental.Customer.Fname = customer.Fname
	rental.Customer.Lname = customer.Lname
	rental.Customer.Username = customer.Username
	rental.Customer.Password = customer.Password
	rental.Customer.UserType = customer.UserType
	rental.Customer.Address = customer.Address
	rental.Customer.Phone = customer.Phone
	rental.Customer.CreatedAt = customer.CreatedAt
	rental.Customer.UpdatedAt = customer.UpdatedAt

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Rental | Created",
		Data:    rental,
	})
}

// GetRentalList godoc
// @Summary     List rental
// @Description get rental
// @Tags        rentals
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Rental}
// @Router      /v1/rental [get]
func (h Handler) GetRentalList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Conf.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Conf.DefaultLimit)
	search := c.DefaultQuery("search", "")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "offset error",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "limit error",
		})
		return
	}
	rentalList, err := h.grpcClients.Rental.GetRentalList(c.Request.Context(), &rental.GetRentalListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    rentalList,
	})
}

// UpdateRental godoc
// @Summary     update rental
// @Description update a new rental
// @Tags        rentals
// @Accept      json
// @Produce     json
// @Param       rental        body     models.UpdateRentalModel true  "rental body"
// @Param       Authorization header   string                   false "Authorization"
// @Success     200           {object} models.JSONResponse{data=models.Rental}
// @Response    400           {object} models.JSONErrorResponse
// @Router      /v1/rental [put]
func (h Handler) UpdateRental(c *gin.Context) {
	var body models.UpdateRentalModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(c.Request.Context(), &car.GetCarByIDRequest{
		Id: body.CarId,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Car not found to create rental",
		})
		return
	}

	customer, err := h.grpcClients.Authorization.GetUserByID(c.Request.Context(), &authorization.GetUserByIDRequest{
		Id: body.CustomerId,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "User is not registired to rental car",
		})
		return
	}

	updated, err := h.grpcClients.Rental.UpdateRental(c.Request.Context(), &rental.UpdateRentalRequest{
		RentalId:   body.RentalId,
		CarId:      car.CarId,
		CustomerId: customer.Id,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
		Payment:    body.Payment,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Faild update!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "rental | Update",
	})
}

// DeleteRental godoc
// @Summary     delete rental by id
// @Description delete an rental by id
// @Tags        rentals
// @Accept      json
// @Param       id            path   string true  "rental ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Rental}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/rental/{id} [delete]
func (h Handler) DeleteRental(c *gin.Context) {
	idStr := c.Param("id")
	deleted, err := h.grpcClients.Rental.DeleteRental(c.Request.Context(), &rental.DeleteRentalRequest{
		RentalId: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Rental have been deleted already!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "rental deleted",
		"data":    deleted,
	})
}
