package handlers

import (
	"car_rental/genprotos/authorization"
	"car_rental/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary     Create user
// @Description create a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user         body     models.CreateUserModel true  "user body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.User}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/user [post]
func (h Handler) CreateUser(c *gin.Context) {
	var body models.CreateUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	user, err := h.grpcClients.Authorization.CreateUser(c.Request.Context(), &authorization.CreateUserRequest{
		Fname:    body.Fname,
		Lname:    body.Lname,
		Username: body.Username,
		Password: body.Password,
		UserType: body.UserType,
		Address:  body.Address,
		Phone:    body.Phone,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "User | Created",
		Data:    user,
	})
}

// GetUserByID godoc
// @Summary     get user by id
// @Description get an user by id
// @Tags        users
// @Accept      json
// @Param       id            path   string true  "User ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.User}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/user/{id} [get]
func (h Handler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation
	user, err := h.grpcClients.Authorization.GetUserByID(c.Request.Context(), &authorization.GetUserByIDRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "User | Get",
		Data:    user,
	})
}

// GetUserList godoc
// @Summary     List user
// @Description get user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.User}
// @Router      /v1/user [get]
func (h Handler) GetUserList(c *gin.Context) {
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
	userList, err := h.grpcClients.Authorization.GetUserList(c.Request.Context(), &authorization.GetUserListRequest{
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
		Data:    userList,
	})
}

// UpdateUser godoc
// @Summary     update user
// @Description update a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user         body     models.UpdateUserModel true  "user body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     200           {object} models.JSONResponse{data=models.User}
// @Response    400           {object} models.JSONErrorResponse
// @Router      /v1/user [put]
func (h Handler) UpdateUser(c *gin.Context) {
	var body models.UpdateUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	updated, err := h.grpcClients.Authorization.UpdateUser(c.Request.Context(), &authorization.UpdateUserRequest{
		Id:       body.Id,
		Fname:    body.Fname,
		Lname:    body.Lname,
		Username: body.Username,
		Password: body.Password,
		UserType: body.UserType,
		Address:  body.Address,
		Phone:    body.Phone,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Faild update!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "user | Update",
	})
}

// DeleteUser godoc
// @Summary     delete user by id
// @Description delete an user by id
// @Tags        users
// @Accept      json
// @Param       id            path   string true  "user ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.User}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/user/{id} [delete]
func (h Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	deleted, err := h.grpcClients.Authorization.DeleteUser(c.Request.Context(), &authorization.DeleteUserRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "User have been deleted already!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
		"data":    deleted,
	})
}
