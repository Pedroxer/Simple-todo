package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload create payload
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token expired or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
