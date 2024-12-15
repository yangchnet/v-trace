package token

import (
	"errors"
	"testing"
	"time"
)

func Test_jwt(t *testing.T) {
	maker, _ := NewJWTMaker(&TokenConfig{
		Secret:     "326c9f73b0f45705a73eb020a32d790a",
		Expiration: time.Millisecond * 100,
	})

	token, _ := maker.CreateToken("aaa", "admin")

	t.Log(token)

	time.Sleep(time.Second)

	_, err := maker.VerifyToken(token)
	if !errors.Is(err, ErrExpiredToken) {
		t.Fail()
	}
}
