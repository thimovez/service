package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/usecase/authorization"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/pkg/validator"
)

type authorizationRoutes struct {
	a authorization.AuthService
	t token.TokenService
	v validator.IValidator
}

func NewAuthorizationRoutes(
	handler *gin.Engine,
	a authorization.AuthService,
	t token.TokenService,
	v validator.IValidator,
) {
	r := &authorizationRoutes{
		a,
		t,
		v,
	}

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/authorization")
	{
		h.POST("/login", r.login)
		h.POST("/registration", r.registration)
	}
}
