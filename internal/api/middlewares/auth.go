package middlewares

import (
	"github.com/thimovez/service/internal/usecase"
	"net/http"
	"strings"
)

type Middleware struct {
	t usecase.TokenService
}

func New(t usecase.TokenService) *Middleware {
	return &Middleware{t}
}

func (m *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Replace this with your actual token validation logic
		claims, err := m.t.VerifyAccessToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		id := claims["userID"].(string)

		err = r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		formData := r.PostForm
		formData.Set("userID", id)

		next(w, r)
	}
}