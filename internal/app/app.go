package app

import (
	"fmt"
	"github.com/thimovez/service/config"
	imageAPI "github.com/thimovez/service/internal/api/image"
	"github.com/thimovez/service/internal/api/middlewares"
	userAPI "github.com/thimovez/service/internal/api/user"
	"github.com/thimovez/service/internal/usecase/image"
	"github.com/thimovez/service/internal/usecase/repo/image-repo"
	"github.com/thimovez/service/internal/usecase/repo/user-repo"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/internal/usecase/user"
	"github.com/thimovez/service/pkg/postgres"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {
	db, err := postgres.SetupDB(cfg.PG.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	userRepo := user_repo.New(db)
	imageRepo := image_repo.New(db)

	tokenCase := token.New()
	userCase := user.New(userRepo, tokenCase)
	imageCase := image.New(imageRepo)

	mux := http.NewServeMux()
	m := middlewares.New(tokenCase)

	userAPI.NewUserRoutes(mux, userCase)
	imageAPI.NewImageRoutes(mux, imageCase, m)

	http.ListenAndServe(cfg.HTTP.Port, mux)
}
