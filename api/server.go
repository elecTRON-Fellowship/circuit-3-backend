package api

import (
	"log"

	db "github.com/elecTRON-Fellowship/formula-1/database/sqlc"
	"github.com/elecTRON-Fellowship/formula-1/pkg/token"
	config "github.com/elecTRON-Fellowship/formula-1/pkg/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app    *fiber.App
	config config.Config
	repo   *db.Repo
	token  token.Token
}

// NewServer is a constructor for Server struct
func NewServer(config config.Config, repo *db.Repo) (*Server, error) {
	// Create a fiber app instance
	token, err := token.NewJWT(config.JWTSecret)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})

	return &Server{
		app,
		config,
		repo,
		token,
	}, nil
}

// Initialize the handlers to respective
// routes
func (s *Server) SetRoutes() error {
	s.app.Use(recover.New())
	s.app.Use(logger.New())

	s.app.Post("/register", s.RegisterUser)
	s.app.Post("/login", s.Login)

	// protected routes
	user := s.app.Group("/user", s.AuthMiddleware())
	user.Get("/", s.GetUserByUserName)
	user.Put("/firstname", s.UpdateFirstName)
	user.Put("/lastname", s.UpdateLastName)
	user.Put("/username", s.UpdateUserName)
	user.Put("/email", s.UpdateEmail)
	user.Put("/password", s.UpdatePassword)
	user.Put("/phone", s.UpdatePhoneNo)
	user.Delete("/", s.DeleteUser)

	s.app.Get("/users", s.AuthMiddleware(), s.ListUsers)

	return nil
}

// StartServer starts the server...
func (s *Server) StartServer(config config.Config) error {
	return s.app.Listen(config.Addr)
}
