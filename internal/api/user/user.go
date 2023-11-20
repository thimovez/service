package user

import (
	"context"
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization"
	"net/http"
)

type userRoutes struct {
	iUserService authorization.AuthUserService
	context      context.Context
}

func NewUserRoutes(handler *http.ServeMux, u authorization.AuthUserService, c context.Context) {
	r := &userRoutes{
		iUserService: u,
		context:      c,
	}

	handler.HandleFunc("/login", r.login)
	handler.HandleFunc("/registration", r.registration)
}

type LoginResponse struct {
	AccessToken string
}

func (u *userRoutes) login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var user entity.UserRequest
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := u.iUserService.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// encode login response to JSON format
	token, err := json.Marshal(LoginResponse{
		AccessToken: accessToken,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(token)
}

func (u *userRoutes) registration(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var user entity.UserRequest
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.iUserService.Registration(user, u.context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
