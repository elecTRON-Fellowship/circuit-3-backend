package token

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the data stored in an access token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	IssuedAT  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// NewPayload is a constructor for Payload
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Payload{
		ID:        tokenID,
		UserName:  username,
		IssuedAT:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}, nil
}

// Valid checks if the token payload is valid
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
