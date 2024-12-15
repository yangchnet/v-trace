package idtools

import "testing"

func Test_NewId(t *testing.T) {
	id := NewId()
	t.Log(id)
}
