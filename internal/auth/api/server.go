package auth

import (
	db "github.com/elecTRON-Fellowship/formula-1/internal/auth/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app  *fiber.App
	repo *db.Repo
}

func NewServer(repo *db.Repo) (*Server, error) {
	// Create a fiber app instance
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})

	app.Use(logger.New())

	return &Server{
		app,
		repo,
	}, nil
}

// Initialize the handlers to respective
// routes
func (s *Server) SetRoutes() error {
	auth := s.app.Group("/api")
	auth.Post("/register", s.RegisterUser)
	auth.Get("/user", s.GetUser)
	return nil
}

func (s *Server) StartServer(addr string) error {
	return s.app.Listen(addr)
}
