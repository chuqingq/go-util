package util

import (
	uuid "github.com/satori/go.uuid"
)

func GenUUids() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}
