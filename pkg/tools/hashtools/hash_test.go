package hashtools

import "testing"

func TestHash(t *testing.T) {
	args := []string{"one", "two", "three"}
	t.Log(Sha256(args))
}
