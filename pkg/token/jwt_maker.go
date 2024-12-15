package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJWTMaker)

// JWTMaker implements the Maker interface.
type JWTMaker struct {
	*TokenConfig
}

type TokenConfig struct {
	Secret     string        `mapstructure:"secret"`
	Expiration time.Duration `mapstructure:"expiration"`
}

const minSecretKeySize = 32

var _ Maker = (*JWTMaker)(nil)

// NewJWTMaker creates a JWTMaker with secret key length = 32.
func NewJWTMaker(conf *TokenConfig) (Maker, error) {
	if len(conf.Secret) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d bytes", minSecretKeySize)
	}

	return &JWTMaker{
		TokenConfig: conf,
	}, nil
}

// CreateToken create a new jwt token with userId and duration.
func (maker *JWTMaker) CreateToken(username, role string) (string, error) {
	payload, err := NewPayload(username, role, maker.Expiration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.Secret))
}

// VerifyToken verify token and return payload.
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// keyFunc check if the token signing algorithm is what you used, if check passed, return the secretKey
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.Secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
