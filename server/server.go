package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	database "github.com/mfulki/ewallet-go-fiber/db"
	"github.com/mfulki/ewallet-go-fiber/handler"
	"github.com/mfulki/ewallet-go-fiber/repository"
	"github.com/mfulki/ewallet-go-fiber/usecase"
)

type Handlers struct {
	AuthHandler *handler.AuthHandler
}

type Server struct {
	db database.DBConnection
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: database.NewDBConnection(db),
	}
}

func (s *Server) SetupServer() *fiber.App {
	userRepository := repository.NewUserRepository(s.db)
	passwordTokenRepository := repository.NewPasswordTokenRepository(s.db)
	walletRepository := repository.NewWalletRepository(s.db)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepository, walletRepository, passwordTokenRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	return SetupRouter(&Handlers{
		AuthHandler: authHandler,
	})
}
