package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayload    = "payload"
)

func (s *Server) AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "authorization header is not provided",
			})
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization type",
			})
		}
		accessToken := fields[1]
		payload, err := s.token.VerifyToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err,
			})
		}
		c.Locals(authorizationPayload, payload.UserName)
		return c.Next()
	}
}
