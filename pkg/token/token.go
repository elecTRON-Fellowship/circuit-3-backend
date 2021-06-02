package token

import "time"

// Token is an interface for managing tokens
type Token interface {
	// CreateToken creates a new token
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken verifies the token provided
	VerifyToken(token string) (*Payload, error)
}
