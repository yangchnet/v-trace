package token

// Maker is a token maker.
//
//go:generate mockgen -source=maker.go -destination=mock_maker.go -package=token . Maker
type Maker interface {
	CreateToken(username, role string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
