package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/thimovez/service/internal/usecase/token"
)

type Auth interface {
	ValidateAuth() gin.HandlerFunc
}

type Middleware struct {
	iTokenService token.TokenService
}

func New(t token.TokenService) *Middleware {
	return &Middleware{t}
}

func (m *Middleware) ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			http.Error(c.Writer, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(c.Writer, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.iTokenService.VerifyAccessToken(token)
		if err != nil {
			http.Error(c.Writer, "Invalid token", http.StatusUnauthorized)
			return
		}

		id := claims["userID"].(string)

		err = c.Request.ParseForm()
		if err != nil {
			http.Error(c.Writer, "Error parsing form data", http.StatusBadRequest)
			return
		}

		formData := c.Request.PostForm
		formData.Set("userID", id)

		c.Next()
	}
}
