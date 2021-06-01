package auth

import (
	"log"
	"strconv"

	db "github.com/elecTRON-Fellowship/formula-1/internal/auth/db/sqlc"
	"github.com/elecTRON-Fellowship/formula-1/pkg/bcrypt"
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

	// hash the user password before storing it
	hashedPass, err := bcrypt.HashPasswd(data.Password)
	if err != nil {
		log.Fatal(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("There was an error, please try again in some time")
	}
	user, err := s.repo.Queries.CreateUser(ctx.Context(), db.CreateUserParams{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UserName:  data.UserName,
		Email:     data.Email,
		Password:  hashedPass,
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
	key := ctx.Params("id")
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

// GetUserByEmail fetches a user's details based on the email provided
func (s *Server) GetUserByEmail(ctx *fiber.Ctx) error {
	key := ctx.Params("email")
	user, err := s.repo.GetUserByEmail(ctx.Context(), key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetUserByUserName fetches a user's details based on the user_name provided
func (s *Server) GetUserByUserName(ctx *fiber.Ctx) error {
	key := ctx.Params("username")
	user, err := s.repo.GetUserByUserName(ctx.Context(), key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// UpdateFirstName updates firstname of the user given his id
func (s *Server) UpdateFirstName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	if err = s.repo.UpdateFirstName(ctx.Context(), db.UpdateFirstNameParams{
		ID:        int64(id),
		FirstName: data.FirstName,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("FirstName updated successfully...")
}

// UpdateLastName updates lastname of the user given his id
func (s *Server) UpdateLastName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	if err = s.repo.UpdateLastName(ctx.Context(), db.UpdateLastNameParams{
		ID:       int64(id),
		LastName: data.LastName,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("LastName updated successfully...")
}

// UpdateUserName updates username of the user given his id
func (s *Server) UpdateUserName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	if err = s.repo.UpdateUserName(ctx.Context(), db.UpdateUserNameParams{
		ID:       int64(id),
		UserName: data.UserName,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("UserName updated successfully...")
}

// UpdateEmail updates email of the user given his id
func (s *Server) UpdateEmail(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	if err = s.repo.UpdateEmail(ctx.Context(), db.UpdateEmailParams{
		ID:    int64(id),
		Email: data.Email,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("Email updated successfully...")
}

// UpdatePassword updates password of the user given his id
func (s *Server) UpdatePassword(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}
	// hash the user password before storing it
	hashedPass, err := bcrypt.HashPasswd(data.Password)
	if err != nil {
		log.Fatal(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("There was an error, please try again in some time")
	}

	if err = s.repo.UpdatePassword(ctx.Context(), db.UpdatePasswordParams{
		ID:       int64(id),
		Password: hashedPass,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("Password updated successfully...")
}

// UpdatePhoneNo updates phone number of the user given his id
func (s *Server) UpdatePhoneNo(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}

	if err = s.repo.UpdatePhoneNo(ctx.Context(), db.UpdatePhoneNoParams{
		ID:      int64(id),
		PhoneNo: data.PhoneNo,
	}); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("Phone Number updated successfully...")
}

// ListUsers returns a list of users given the limit and offset in the url query
func (s *Server) ListUsers(ctx *fiber.Ctx) error {
	// get the value of limit
	limitString := ctx.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}
	// get the value of offset
	offsetString := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}
	users, err := s.repo.ListUsers(ctx.Context(), db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// DeleteUser deletes a user from the db
func (s *Server) DeleteUser(ctx *fiber.Ctx) error {
	// get the id from query params
	key := ctx.Params("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
		return fiber.ErrBadRequest
	}
	if err := s.repo.DeleteUser(ctx.Context(), int64(id)); err != nil {
		log.Fatal(err)
		return fiber.ErrInternalServerError
	}
	return ctx.Status(fiber.StatusOK).JSON("User successfully deleted...")
}
