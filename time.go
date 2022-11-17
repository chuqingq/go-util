package util

import (
	"time"
	"github.com/satori/go.uuid"
)

func GetMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func GenUUids() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}