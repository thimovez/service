package app

import (
	"fmt"
	"github.com/pressly/goose"
	"github.com/thimovez/service/config"
	imageAPI "github.com/thimovez/service/internal/api/image"
	"github.com/thimovez/service/internal/api/middlewares"
	userAPI "github.com/thimovez/service/internal/api/user"
	"github.com/thimovez/service/internal/providers/auth"
	"github.com/thimovez/service/internal/providers/helpers"
	"github.com/thimovez/service/internal/usecase/image"
	imageRepo "github.com/thimovez/service/internal/usecase/repo/postgres/image"
	userRepo "github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/internal/usecase/user"
	"github.com/thimovez/service/pkg/postgres"
	"log"
	"net/http"
	"time"
)

const tokenTime = 12

func Run(cfg *config.Config) {
	db, err := postgres.SetupDB(cfg.PG.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	err = goose.Up(db, "./migrations")
	if err != nil {
		log.Fatal(fmt.Errorf("migration error: %w", err))
	}

	userRepoPG := userRepo.New(db)
	imageRepoPG := imageRepo.New(db)

	expiration := time.Now().Add(time.Hour * tokenTime)
	jwtProvider, err := auth.NewJWTProvider(cfg.TOKEN.Secret, expiration)
	if err != nil {
		fmt.Printf("Error initializing JWT provider: %v\n", err)
		return
	}
	helperProvider := helpers.NewHelperProvider()

	tokenUseCase := token.New(jwtProvider)
	userUseCase := user.New(userRepoPG, tokenUseCase, helperProvider)
	imageUseCase := image.New(imageRepoPG, helperProvider)

	mux := http.NewServeMux()
	m := middlewares.New(tokenUseCase)

	userAPI.NewUserRoutes(mux, userUseCase)
	imageAPI.NewImageRoutes(mux, imageUseCase, m)

	http.ListenAndServe(cfg.HTTP.Port, mux)
}
