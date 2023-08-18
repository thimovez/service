package http1

import (
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"log"
	"net/http"
)

type userRoutes struct {
	t usecase.User
}

func NewUserRoutes(handler *http.ServeMux, t usecase.User) {
	r := &userRoutes{t}

	handler.HandleFunc("/login", r.login)
}

func (r *userRoutes) login(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var user entity.UserRequest
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	err = r.t.Login(user.Username, user.Password)
	if err != nil {
		log.Fatalf("login service error %s", err)
	}
}
