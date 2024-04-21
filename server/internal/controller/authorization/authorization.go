package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization"
	"net/http"
)

type authorizationRoutes struct {
	a authorization.AuthUserService
}

func newAuthorizationRoutes(handler *gin.RouterGroup, a authorization.AuthUserService) {
	r := &authorizationRoutes{a}

	h := handler.Group("/authorization")
	{
		h.POST("/login", r.login)
		h.POST("/registration", r.registration)
	}
}

func (r *authorizationRoutes) login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user entity.AuthorizationReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.a.VerifyLoginData(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *authorizationRoutes) registration(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user entity.UserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.a.VerifyRegistrationData(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}