package user

import (
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"net/http"
)

type userRoutes struct {
	u usecase.UserService
}

func NewUserRoutes(handler *http.ServeMux, u usecase.UserService) {
	r := &userRoutes{u}

	handler.HandleFunc("/login", r.login)
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

	accessToken, err := u.u.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	marshal, err := json.Marshal(entity.LoginResponse{
		AccessToken: accessToken,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
