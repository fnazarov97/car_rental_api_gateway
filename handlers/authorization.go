package handlers

import (
	"car_rental/genprotos/authorization"
	"car_rental/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
func (h Handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasAccessResponse, err := h.grpcClients.Authorization.HasAccess(c.Request.Context(), &authorization.TokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if !hasAccessResponse.HasAccess {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if userType != "*" {
			if hasAccessResponse.User.UserType != userType {
				c.JSON(http.StatusUnauthorized, "Permission Denied")
				c.Abort()
			}
		}

		c.Set("auth_username", hasAccessResponse.User.Username)
		c.Set("auth_user_id", hasAccessResponse.User.Id)

		c.Next()
		//
	}
}

// Login godoc
// @Summary     Login
// @Description Login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       login body     models.LoginModel true "Login body"
// @Success     201   {object} models.JSONResponse{data=models.TokenResponse}
// @Failure     400   {object} models.JSONErrorResponse
// @Router      /v1/login [post]
func (h Handler) Login(c *gin.Context) {
	var body models.LoginModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "err.Error()1"})
		return
	}
	log.Println("-------------------------->", body.Password, body.Username)
	// TODO - validation should be here
	tokenResponse, err := h.grpcClients.Authorization.Login(c.Request.Context(), &authorization.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "err.Error()2",
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | GetList",
		Data:    tokenResponse,
	})
}
