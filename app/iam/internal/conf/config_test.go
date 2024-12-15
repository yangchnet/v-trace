package conf

import "testing"

func Test_load(t *testing.T) {
	c := LoadConfig("../../config")
	t.Log(c)
}
