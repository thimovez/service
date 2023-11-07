package user

import (
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"net/http"
)

type userRoutes struct {
	iUserService usecase.UserService
}

func NewUserRoutes(handler *http.ServeMux, u usecase.UserService) {
	r := &userRoutes{u}

	handler.HandleFunc("/login", r.login)
	handler.HandleFunc("/registration", r.registration)
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
	token, err := json.Marshal(entity.LoginResponse{
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

	err = u.iUserService.Registration(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
