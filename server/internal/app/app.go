package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
	"github.com/thimovez/service/config"
	userAPI "github.com/thimovez/service/internal/controller/user"
	"github.com/thimovez/service/internal/providers/bcrypt"
	"github.com/thimovez/service/internal/providers/uuid"
	"github.com/thimovez/service/internal/usecase/authorization"
	userRepo "github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/internal/usecase/token/tokenapi"
	"github.com/thimovez/service/pkg/httpserver"
	"github.com/thimovez/service/pkg/logger"
	"github.com/thimovez/service/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const tokenTime = 12

func Run(cfg *config.Config) {
	l := logger.New(cfg.LOG.Level)

	db, err := postgres.SetupDB(cfg.PG.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	err = goose.Up(db, "../migrations")
	if err != nil {
		log.Fatal(fmt.Errorf("migration error: %w", err))
	}

	userRepoPG := userRepo.New(db)
	//imageRepoPG := imageRepo.New(db)
	//commentRepoPG := commentRepo.New(db)

	expiration := time.Now().Add(time.Hour * tokenTime)
	jwtProvider, err := tokenapi.NewJWTProvider(cfg.TOKEN.Secret, expiration)
	if err != nil {
		fmt.Printf("Error initializing JWT provider: %v\n", err)
		return
	}

	UUIDProvider := uuid.NewUUIDProvider()
	bcryptProvider := bcrypt.NewBcryptProvider()

	tokenUseCase := token.New(jwtProvider)
	userUseCase := authorization.New(userRepoPG, tokenUseCase, UUIDProvider, bcryptProvider)
	//imageUseCase := image.New(imageRepoPG, UUIDProvider)
	//commentUseCase := comment.New(commentRepoPG)

	handler := gin.New()
	//mux := http.NewServeMux()
	//m := middlewares.New(tokenUseCase)

	userAPI.NewRouter(handler, userUseCase)
	//imageAPI.NewImageRoutes(mux, imageUseCase, m)
	//commentAPI.NewCommentRoutes(mux, commentUseCase)

	httpServer := httpserver.New(handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}

}
