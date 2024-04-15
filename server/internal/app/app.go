package app

import (
	"context"
	"fmt"
	"github.com/pressly/goose"
	"github.com/thimovez/service/config"
	commentAPI "github.com/thimovez/service/internal/api/comment"
	imageAPI "github.com/thimovez/service/internal/api/image"
	"github.com/thimovez/service/internal/api/middlewares"
	userAPI "github.com/thimovez/service/internal/api/user"
	"github.com/thimovez/service/internal/providers/auth"
	"github.com/thimovez/service/internal/providers/bcrypt"
	"github.com/thimovez/service/internal/providers/uuid"
	"github.com/thimovez/service/internal/usecase/authorization"
	"github.com/thimovez/service/internal/usecase/comment"
	"github.com/thimovez/service/internal/usecase/image"
	commentRepo "github.com/thimovez/service/internal/usecase/repo/postgres/comment"
	imageRepo "github.com/thimovez/service/internal/usecase/repo/postgres/image"
	userRepo "github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userRepoPG := userRepo.New(db)
	imageRepoPG := imageRepo.New(db)
	commentRepoPG := commentRepo.New(db)

	expiration := time.Now().Add(time.Hour * tokenTime)
	jwtProvider, err := auth.NewJWTProvider(cfg.TOKEN.Secret, expiration)
	if err != nil {
		fmt.Printf("Error initializing JWT provider: %v\n", err)
		return
	}
	UUIDProvider := uuid.NewUUIDProvider()
	bcryptProvider := bcrypt.NewBcryptProvider()

	tokenUseCase := token.New(jwtProvider)
	userUseCase := authorization.New(userRepoPG, tokenUseCase, UUIDProvider, bcryptProvider)
	imageUseCase := image.New(imageRepoPG, UUIDProvider)
	commentUseCase := comment.New(commentRepoPG)

	mux := http.NewServeMux()
	m := middlewares.New(tokenUseCase)

	userAPI.NewUserRoutes(mux, userUseCase, ctx)
	imageAPI.NewImageRoutes(mux, imageUseCase, m)
	commentAPI.NewCommentRoutes(mux, commentUseCase)

	http.ListenAndServe(cfg.HTTP.Port, mux)
}
