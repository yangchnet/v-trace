package idtools

import (
	"strings"

	"github.com/google/uuid"
)

func NewId() string {
	uid := uuid.New().String()
	list := strings.Split(uid, "-")
	return list[len(list)-1]
}
