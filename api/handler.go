package api

import (
	"database/sql"
	"log"
	"strconv"

	db "github.com/elecTRON-Fellowship/formula-1/database/sqlc"
	"github.com/elecTRON-Fellowship/formula-1/pkg/bcrypt"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhoneNo   string `json:"phone_no"`
}

// RegisterUser creates a new user entry in the db
func (s *Server) RegisterUser(ctx *fiber.Ctx) error {
	data := new(user)
	// parse the request body elements and store them in data
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data format",
		})
	}

	// hash the user password before storing it
	hashedPass, err := bcrypt.HashPasswd(data.Password)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error creating account, please try again after sometime...",
		})
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
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error creating account, please try again after sometime...",
		})
	}
	accessToken, err := s.token.CreateToken(user.UserName, s.config.JWTDuration)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error loging you in. Please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": accessToken,
		"data":        user,
		"message":     "User successfully added...",
	})
}

// GetUser fetches a user's data given his id
func (s *Server) GetUser(ctx *fiber.Ctx) error {
	// get the id from query params
	key := ctx.Query("id")
	// convert the id from string to int64
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error retreiving account, please try again after sometime...",
		})
	}

	user, err := s.repo.GetUser(ctx.Context(), int64(id))
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error retreiving account, please try again after sometime...",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"message": "User successfully fetched",
	})
}

// GetUserByEmail fetches a user's details based on the email provided
func (s *Server) GetUserByEmail(ctx *fiber.Ctx) error {
	key := ctx.Query("email")
	user, err := s.repo.GetUserByEmail(ctx.Context(), key)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error retreiving account, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"message": "User successfully fetched",
	})
}

// GetUserByUserName fetches a user's details based on the user_name provided
func (s *Server) GetUserByUserName(ctx *fiber.Ctx) error {
	//username := ctx.Get(authorizationPayload)
	username := ctx.Locals(authorizationPayload)
	user, err := s.repo.GetUserByUserName(ctx.Context(), username.(string))
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error retreiving account, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"message": "User successfully fetched",
	})
}

// UpdateFirstName updates firstname of the user given his id
func (s *Server) UpdateFirstName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}
	username := ctx.Locals(authorizationPayload)

	if err := s.repo.UpdateFirstName(ctx.Context(), db.UpdateFirstNameParams{
		UserName:  username.(string),
		FirstName: data.FirstName,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating firstname, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Firstname successfully updated",
	})
}

// UpdateLastName updates lastname of the user given his id
func (s *Server) UpdateLastName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}

	username := ctx.Locals(authorizationPayload)
	if err := s.repo.UpdateLastName(ctx.Context(), db.UpdateLastNameParams{
		UserName: username.(string),
		LastName: data.LastName,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating lastname, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Lastname successfully updated",
	})
}

// UpdateUserName updates username of the user given his id
func (s *Server) UpdateUserName(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}
	username := ctx.Locals(authorizationPayload)

	if err := s.repo.UpdateUserName(ctx.Context(), db.UpdateUserNameParams{
		UserName:   username.(string),
		UserName_2: data.UserName,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating username, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Username successfully updated",
	})
}

// UpdateEmail updates email of the user given his id
func (s *Server) UpdateEmail(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}

	username := ctx.Locals(authorizationPayload)
	if err := s.repo.UpdateEmail(ctx.Context(), db.UpdateEmailParams{
		UserName: username.(string),
		Email:    data.Email,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating email, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Email successfully updated",
	})
}

// UpdatePassword updates password of the user given his id
func (s *Server) UpdatePassword(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}

	username := ctx.Locals(authorizationPayload)
	// hash the user password before storing it
	hashedPass, err := bcrypt.HashPasswd(data.Password)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error, please try again after sometime...",
		})
	}

	if err = s.repo.UpdatePassword(ctx.Context(), db.UpdatePasswordParams{
		UserName: username.(string),
		Password: hashedPass,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating password, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password successfully updated",
	})
}

// UpdatePhoneNo updates phone number of the user given his id
func (s *Server) UpdatePhoneNo(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}

	username := ctx.Locals(authorizationPayload)
	if err := s.repo.UpdatePhoneNo(ctx.Context(), db.UpdatePhoneNoParams{
		UserName: username.(string),
		PhoneNo:  data.PhoneNo,
	}); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error updating phone number, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Phone number successfully updated",
	})
}

// ListUsers returns a list of users given the limit and offset in the url query
func (s *Server) ListUsers(ctx *fiber.Ctx) error {
	// get the value of limit
	limitString := ctx.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}
	// get the value of offset
	offsetString := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corrupted data sent",
		})
	}
	users, err := s.repo.ListUsers(ctx.Context(), db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error fetching accounts, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "Users successfully fetched.",
	})
}

// DeleteUser deletes a user from the db
func (s *Server) DeleteUser(ctx *fiber.Ctx) error {
	username := ctx.Locals(authorizationPayload)
	if err := s.repo.DeleteUser(ctx.Context(), username.(string)); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error deleting the account, please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User successfully deleted.",
	})
}

// Login - it should be self explanatory
func (s *Server) Login(ctx *fiber.Ctx) error {
	data := new(user)
	if err := ctx.BodyParser(&data); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data, please try again!",
		})
	}
	if data.PhoneNo == "" {
		user, err := s.repo.GetUserByUserName(ctx.Context(), data.UserName)
		if err != nil {
			log.Print(err)
			if err == sql.ErrNoRows {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "The username provided does not exist, maybe try registering first if you haven't.",
				})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "There was an error fetching the user, please try again after sometime...",
			})
		}
		if err = bcrypt.VerifyPasswd(user.Password, data.Password); err != nil {
			log.Print(err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "The password doesn't match with the username provided",
			})
		}
		accessToken, err := s.token.CreateToken(user.UserName, s.config.JWTDuration)
		if err != nil {
			log.Print(err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "There was an error loging you in. Please try again after sometime...",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"accessToken": accessToken,
			"data":        user,
			"message":     "User successfully logged in.",
		})
	}
	user, err := s.repo.GetUserByPhoneNo(ctx.Context(), data.PhoneNo)
	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "The username provided does not exist, maybe try registering first if you haven't.",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error fetching the user, please try again after sometime...",
		})
	}
	if err = bcrypt.VerifyPasswd(user.Password, data.Password); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "The password doesn't match with the username provided",
		})
	}
	accessToken, err := s.token.CreateToken(user.UserName, s.config.JWTDuration)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "There was an error loging you in. Please try again after sometime...",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": accessToken,
		"data":        user,
		"message":     "User successfully logged in.",
	})
}
