package token

import (
	"time"
)

// Payload is Jwt payload.
type Payload struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload create a new token payload.
func NewPayload(username, role string, duration time.Duration) (*Payload, error) {
	// tokenID, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, err
	// }

	return &Payload{
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

// Valid validates the payload is or not valid.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

// func GetPayload(c *gin.Context) (*Payload, error) {
// 	if payload, ok := c.Get("authorization_payload"); ok {
// 		return payload.(*Payload), nil
// 	}

// 	return nil, fmt.Errorf("解析token失败")
// }
