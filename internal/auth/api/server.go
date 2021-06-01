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
	s.app.Post("/register", s.RegisterUser)
	s.app.Get("/users", s.ListUsers)

	user := s.app.Group("/user")
	user.Get("/id/:id", s.GetUser)
	user.Get("/email/:email", s.GetUserByEmail)
	user.Get("/username/:username", s.GetUserByUserName)
	user.Put("/:id/firstname", s.UpdateFirstName)
	user.Put("/:id/lastname", s.UpdateLastName)
	user.Put("/:id/username", s.UpdateUserName)
	user.Put("/:id/email", s.UpdateEmail)
	user.Put("/:id/password", s.UpdatePassword)
	user.Put("/:id/phone", s.UpdatePhoneNo)
	user.Delete("/:id", s.DeleteUser)

	s.app.Post("/login", s.Login)

	return nil
}

func (s *Server) StartServer(addr string) error {
	return s.app.Listen(addr)
}
