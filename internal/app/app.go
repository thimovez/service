package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
	"github.com/thimovez/service/config"
	userAPI "github.com/thimovez/service/internal/controller/authorization"
	"github.com/thimovez/service/internal/usecase/authorization"
	"github.com/thimovez/service/internal/usecase/authorization/bcryptapi"
	"github.com/thimovez/service/internal/usecase/authorization/uuidapi"
	userRepo "github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
	"github.com/thimovez/service/internal/usecase/token/tokenapi"
	"github.com/thimovez/service/pkg/httpserver"
	"github.com/thimovez/service/pkg/logger"
	"github.com/thimovez/service/pkg/postgres"
	"github.com/thimovez/service/pkg/validator"
)

const accessTime = 2
const refreshTime = 12

func Run(cfg *config.Config) {
	l := logger.New(cfg.LOG.Level)

	// Construct the connection string
	connStr := fmt.Sprintf("user=%s password=%s port=%s dbname=%s sslmode=disable",
		cfg.PG.Username,
		cfg.PG.Password,
		cfg.PG.Port,
		cfg.PG.Database)

	db, err := postgres.SetupDB(connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	err = goose.Up(db, "./migrations")
	if err != nil {
		log.Fatal(fmt.Errorf("migration error: %w", err))
	}

	jwtProvider, err := tokenapi.NewJWTProvider(cfg.TOKEN.Secret)
	if err != nil {
		fmt.Printf("Error initializing JWT provider: %v\n", err)
		return
	}

	AccessExp := time.Now().Add(time.Hour * accessTime)
	RefreshExp := time.Now().Add(time.Hour * accessTime)
	secret = cfg.TOKEN.Secret

	userUseCase := authorization.New(
		userRepo.New(db),
		uuidapi.NewUUIDProvider(),
		bcryptapi.NewBcryptProvider(),
	)

	//imageRepoPG := imageRepo.New(db)
	//commentRepoPG := commentRepo.New(db)
	//imageUseCase := image.New(imageRepoPG, UUIDProvider)
	//commentUseCase := comment.New(commentRepoPG)

	handler := gin.New()
	//mux := http.NewServeMux()
	//m := middlewares.New(
	//	token.New(jwtProvider),
	//)

	t := token.New(jwtProvider, AccessExp, RefreshExp, secret)
	userAPI.NewAuthorizationRoutes(
		handler,
		userUseCase,
		t,
		validator.New(),
	)
	//imageAPI.NewImageRoutes(mux, imageUseCase, m)
	//commentAPI.NewCommentRoutes(mux, commentUseCase)

	gin.SetMode(gin.DebugMode)
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
