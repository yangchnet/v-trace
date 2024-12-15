package hashtools

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Sha256(args ...any) string {
	h := sha256.New()

	data, err := json.Marshal(args)
	if err != nil {
		data = []byte{}
	}

	h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}
