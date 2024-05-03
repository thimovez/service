package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/pkg/validator"
	"net/http"
)

type authorizationRoutes struct {
	a authorization.AuthUserService
	t token.TokenService
	v validator.IValidator
}

func newAuthorizationRoutes(
	handler *gin.RouterGroup,
	a authorization.AuthUserService,
	t token.TokenService,
	v validator.IValidator,
) {
	r := &authorizationRoutes{
		a,
		t,
		v,
	}

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

	c.JSON(http.StatusOK, res)
}

func (r *authorizationRoutes) registration(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user entity.UserRegistrationReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.v.ValidateStruct(user)
	if err != nil {
		//fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	err = r.a.VerifyRegistrationData(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
