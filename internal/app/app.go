package app

import (
	"fmt"
	"github.com/thimovez/service/config"
	"github.com/thimovez/service/internal/controller/http1"
	"github.com/thimovez/service/internal/usecase"
	"github.com/thimovez/service/internal/usecase/repo"
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

	userRepo := repo.New(db)

	userCase := usecase.New(userRepo)

	mux := http.NewServeMux()

	http1.NewUserRoutes(mux, userCase)

	http.ListenAndServe(cfg.HTTP.Port, mux)
}
