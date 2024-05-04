package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/entity"
	"net/http"
)

func (r *authorizationRoutes) login(c *gin.Context) {
	var user entity.AuthorizationReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validData, err := r.a.VerifyLoginData(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := r.t.GenerateAccessToken(validData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := r.t.GenerateRefreshToken(validData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"refreshToken",
		refreshToken,
		36000,
		"/",
		"localhost",
		false,
		true,
	)

	res := entity.AuthorizationRes{
		User: entity.UserRes{
			ID:       validData.User.ID,
			Username: validData.User.Username,
		},
		Token: entity.Token{
			AccessToken: accessToken,
		},
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, res)
}

func (r *authorizationRoutes) registration(c *gin.Context) {
	var user entity.UserRegistrationReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors := r.v.ValidateStruct(user)
	if errors != nil {
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	err = r.a.VerifyRegistrationData(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
