package auth

import (
	"log"
	"strconv"

	db "github.com/elecTRON-Fellowship/formula-1/internal/auth/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhoneNo   int32  `json:"phone_no"`
}

// RegisterUser creates a new user entry in the db
func (s *Server) RegisterUser(ctx *fiber.Ctx) error {
	data := new(user)
	// parse the request body elements and store them in data
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	user, err := s.repo.Queries.CreateUser(ctx.Context(), db.CreateUserParams{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UserName:  data.UserName,
		Email:     data.Email,
		Password:  data.Password,
		PhoneNo:   data.PhoneNo,
	})
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetUser fetches a user's data given his id
func (s *Server) GetUser(ctx *fiber.Ctx) error {
	// get the id from query params
	key := ctx.Query("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}

	user, err := s.repo.GetUser(ctx.Context(), int64(id))
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
