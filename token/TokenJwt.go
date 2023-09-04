package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTMaker struct {
	secretKey string
}

func NewJwtToken(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < 32 {
		return nil, errors.New("invalid key size: size must be 32")
	}
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}
func (jwtM *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//hMac := hmac.New(sha256.New, []byte(jwtM.secretKey))
	tokenString, err := token.SignedString([]byte(jwtM.secretKey))
	if err != nil {
		log.Println(err)
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
