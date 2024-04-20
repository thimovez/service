package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization"
	"net/http"
)

type authorizationRoutes struct {
	a *authorization.AuthUserUseCase
}

func newAuthorizationRoutes(handler *gin.RouterGroup, a *authorization.AuthUserUseCase) {
	r := &authorizationRoutes{a}

	h := handler.Group("/authorization")
	{
		h.POST("/login", r.login)
		h.POST("/registration", r.registration)
	}
}

type LoginResponse struct {
	Tokens struct {
		AccessToken string `json:"access_token"`
	} `json:"tokens"`
}

func (r *authorizationRoutes) login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user entity.UserRequest
	var res = LoginResponse{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := r.a.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res.Tokens.AccessToken = accessToken

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

	err = r.a.Registration(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
