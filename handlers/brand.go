package handlers

import (
	"car_rental/genprotos/brand"
	"car_rental/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBrand godoc
// @Summary     Create brand
// @Description create a new brand
// @Tags        brands
// @Accept      json
// @Produce     json
// @Param       brand         body     models.CreateBrandModel true  "brand body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Brand}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/brand [post]
func (h Handler) CreateBrand(c *gin.Context) {
	var body models.CreateBrandModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	brand, err := h.grpcClients.Brand.CreateBrand(c.Request.Context(), &brand.CreateBrandRequest{
		Name:        body.Name,
		Discription: body.Discription,
		Year:        body.Year,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Brand | Created",
		Data:    brand,
	})
}

// GetBrandByID godoc
// @Summary     get brand by id
// @Description get an brand by id
// @Tags        brands
// @Accept      json
// @Param       id            path   string true  "Brand ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Brand}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/brand/{id} [get]
func (h Handler) GetBrandByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation
	brand, err := h.grpcClients.Brand.GetBrandByID(c.Request.Context(), &brand.GetBrandByIDRequest{
		BrandId: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Brand | Created",
		Data:    brand,
	})
}

// GetBrandList godoc
// @Summary     List brand
// @Description get brand
// @Tags        brands
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Brand}
// @Router      /v1/brand [get]
func (h Handler) GetBrandList(c *gin.Context) {
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
	brandList, err := h.grpcClients.Brand.GetBrandList(c.Request.Context(), &brand.GetBrandListRequest{
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
		Data:    brandList,
	})
}

// UpdateBrand godoc
// @Summary     update brand
// @Description update a new brand
// @Tags        brands
// @Accept      json
// @Produce     json
// @Param       brand         body     models.UpdateBrandModel true  "brand body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     200           {object} models.JSONResponse{data=models.Brand}
// @Response    400           {object} models.JSONErrorResponse
// @Router      /v1/brand [put]
func (h Handler) UpdateBrand(c *gin.Context) {
	var body models.UpdateBrandModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	updated, err := h.grpcClients.Brand.UpdateBrand(c.Request.Context(), &brand.UpdateBrandRequest{
		Id:          body.BrandId,
		Name:        body.Name,
		Discription: body.Discription,
		Year:        body.Year,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Faild update!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "brand | Update",
	})
}

// DeleteBrand godoc
// @Summary     delete brand by id
// @Description delete an brand by id
// @Tags        brands
// @Accept      json
// @Param       id            path   string true  "brand ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Brand}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/brand/{id} [delete]
func (h Handler) DeleteBrand(c *gin.Context) {
	idStr := c.Param("id")
	deleted, err := h.grpcClients.Brand.DeleteBrand(c.Request.Context(), &brand.DeleteBrandRequest{
		BrandId: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Brand have been deleted already!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "brand deleted",
		"data":    deleted,
	})
}
