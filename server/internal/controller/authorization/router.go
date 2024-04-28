package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/usecase/authorization"
	"github.com/thimovez/service/internal/usecase/token"
)

func NewRouter(handler *gin.Engine, a *authorization.AuthUserUseCase, t token.TokenService) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{

		newAuthorizationRoutes(h, a, t)
	}
}
