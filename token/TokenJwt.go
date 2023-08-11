package token

import (
	"errors"
	"fmt"

	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/util"
	"github.com/golang-jwt/jwt"
)

type JWTMaker struct {
	secretKey string
}

func NewJwtToken(secretKey string) *JWTMaker {
	return &JWTMaker{
		secretKey: secretKey,
	}
}
func (jwtM *JWTMaker) CreateToken(u sqlc.User, conf util.Config) (string, error) {
	payload, err := NewPayload(u.Username, conf.TokenDuration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)

	tokenString, err := token.SignedString(jwtM.secretKey)
	if err != nil {
		return "", fmt.Errorf("Error generating token:%s", err)
	}
	return tokenString, nil
}

func (jwtM *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken // check signing method
		}
		return []byte(jwtM.secretKey), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
