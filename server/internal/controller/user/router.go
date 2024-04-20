package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/service/internal/usecase/authorization"
)

func NewRouter(handler *gin.Engine, a *authorization.AuthUserUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{

		newAuthorizationRoutes(h, a)
	}
}
